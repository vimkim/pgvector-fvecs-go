package main

import (
	"context"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pgvector/pgvector-go"
)

type VectorData struct {
	Vec pgvector.Vector
}

func main() {
	// Command-line flags
	dbname := flag.String("dbname", "", "Database name (required)")
	host := flag.String("host", "localhost", "Database host")
	port := flag.String("port", "5432", "Database port")
	user := flag.String("user", "", "Database user (required)")
	password := flag.String("password", "", "Database password")
	tablename := flag.String("tablename", "items", "Name of the table to load vectors into")
	fvecsFile := flag.String("fvecs", "", "Path to fvecs file (required)")
	workers := flag.Int("workers", 4, "Number of concurrent workers for insertion")
	flag.Parse()

	// Check required flags
	if *dbname == "" || *user == "" || *fvecsFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Build the connection string
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", *user, *password, *host, *port, *dbname)
	ctx := context.Background()

	// Use a connection pool
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}
	defer pool.Close()

	// Enable the pgvector extension
	_, err = pool.Exec(ctx, "CREATE EXTENSION IF NOT EXISTS vector")
	if err != nil {
		log.Fatalf("Error enabling pgvector extension: %v", err)
	}

	// Open the fvecs file
	file, err := os.Open(*fvecsFile)
	if err != nil {
		log.Fatalf("Failed to open fvecs file: %v", err)
	}
	defer file.Close()

	// Read the dimension of the first vector for table definition.
	var dim int32
	if err := binary.Read(file, binary.LittleEndian, &dim); err != nil {
		log.Fatalf("Error reading dimension from fvecs file: %v", err)
	}
	// Reset file pointer back to start.
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		log.Fatalf("Error resetting file pointer: %v", err)
	}

	// Drop table if exists
	dropTableSQL := fmt.Sprintf(`DROP TABLE IF EXISTS %s;`, *tablename)
	if _, err := pool.Exec(ctx, dropTableSQL); err != nil {
		log.Fatalf("Error dropping table %s: %v", *tablename, err)
	}

	// Create the table if it does not exist.
	createTableSQL := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id int,
			vec vector(%d)
		)
	`, *tablename, dim)
	if _, err := pool.Exec(ctx, createTableSQL); err != nil {
		log.Fatalf("Error creating table %s: %v", *tablename, err)
	}

	// Check if the table is empty by querying the row count
	checkEmptySQL := fmt.Sprintf(`SELECT count(*) FROM %s`, *tablename)
	var count int
	err = pool.QueryRow(ctx, checkEmptySQL).Scan(&count)
	if err != nil {
		log.Fatalf("Error querying the number of rows in table %s: %v", *tablename, err)
	}

	if count > 0 {
		log.Fatalf("Error: %d rows already exist in table %s", count, *tablename)
	}

	// Create a channel to queue vectors.
	vectorCh := make(chan pgvector.Vector, 10000)
	var wg sync.WaitGroup

	// Worker function to insert vectors.
	workerFunc := func(workerID int) {
		defer wg.Done()
		for vector := range vectorCh {
			insertSQL := fmt.Sprintf("INSERT INTO %s (vec) VALUES ($1)", *tablename)
			// Each worker gets a connection from the pool.
			if _, err := pool.Exec(ctx, insertSQL, vector); err != nil {
				log.Printf("Worker %d: Error inserting vector: %v", workerID, err)
			}
		}
	}

	// Start worker goroutines.
	for i := 0; i < *workers; i++ {
		wg.Add(1)
		go workerFunc(i)
	}

	// Read vectors from file and send them to the channel.
	inserted := 0
	for {
		var currentDim int32
		err := binary.Read(file, binary.LittleEndian, &currentDim)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading vector dimension: %v", err)
		}
		if currentDim != dim {
			log.Fatalf("Dimension mismatch: expected %d but got %d", dim, currentDim)
		}

		vec := make([]float32, currentDim)
		if err := binary.Read(file, binary.LittleEndian, &vec); err != nil {
			log.Fatalf("Error reading vector values: %v", err)
		}
		vectorCh <- pgvector.NewVector(vec)

		inserted++
		if inserted%10000 == 0 {
			fmt.Printf("%d vectors queued for insertion...\n", inserted)
		}
	}

	// Close the channel to signal workers no more vectors are coming.
	close(vectorCh)
	// Wait for all workers to finish.
	wg.Wait()

	fmt.Printf("Finished inserting %d vectors into table '%s'.\n", inserted, *tablename)
}
