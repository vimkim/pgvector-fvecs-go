package main

import (
	"context"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/pgvector/pgvector-go"
	pgxvec "github.com/pgvector/pgvector-go/pgx"
)

func main() {
	// Command-line flags
	dbname := flag.String("dbname", "", "Database name (required)")
	host := flag.String("host", "localhost", "Database host")
	port := flag.String("port", "5432", "Database port")
	user := flag.String("user", "", "Database user (required)")
	password := flag.String("password", "", "Database password")
	tablename := flag.String("tablename", "items", "Name of the table to load vectors into")
	fvecsFile := flag.String("fvecs", "", "Path to fvecs file (required)")
	flag.Parse()

	// Check required flags
	if *dbname == "" || *user == "" || *fvecsFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Build the connection string
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", *user, *password, *host, *port, *dbname)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close(ctx)

	// Enable the pgvector extension
	_, err = conn.Exec(ctx, "CREATE EXTENSION IF NOT EXISTS vector")
	if err != nil {
		log.Fatalf("Error enabling pgvector extension: %v", err)
	}

	// Register pgvector types with the connection (now using pgx v5)
	err = pgxvec.RegisterTypes(ctx, conn)
	if err != nil {
		log.Fatalf("Error registering pgvector types: %v", err)
	}

	// Open the fvecs file
	file, err := os.Open(*fvecsFile)
	if err != nil {
		log.Fatalf("Failed to open fvecs file: %v", err)
	}
	defer file.Close()

	// Read the dimension of the first vector to use for table definition.
	// fvecs files start with an int32 indicating the vector dimension.
	var dim int32
	if err := binary.Read(file, binary.LittleEndian, &dim); err != nil {
		log.Fatalf("Error reading dimension from fvecs file: %v", err)
	}
	// Reset file pointer back to the start of the file so all vectors are read.
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		log.Fatalf("Error resetting file pointer: %v", err)
	}

	// Create the table if it does not exist.
	// The table will have an id column and an embedding column of type vector(dim).
	createTableSQL := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
            id int,
			embedding vector(%d)
		)
	`, *tablename, dim)
	if _, err := conn.Exec(ctx, createTableSQL); err != nil {
		log.Fatalf("Error creating table %s: %v", *tablename, err)
	}

	// Begin reading vectors from the file and inserting them.
	inserted := 0
	for {
		// Read the dimension for this vector record.
		var currentDim int32
		err := binary.Read(file, binary.LittleEndian, &currentDim)
		if err == io.EOF {
			// End of file reached
			break
		}
		if err != nil {
			log.Fatalf("Error reading vector dimension: %v", err)
		}
		// Ensure the vector has the same dimension as the first one.
		if currentDim != dim {
			log.Fatalf("Dimension mismatch: expected %d but got %d", dim, currentDim)
		}

		// Read the vector values (each as float32)
		vec := make([]float32, currentDim)
		if err := binary.Read(file, binary.LittleEndian, &vec); err != nil {
			log.Fatalf("Error reading vector values: %v", err)
		}

		// Create a pgvector instance from the slice
		vector := pgvector.NewVector(vec)

		// Insert the vector into the table.
		insertSQL := fmt.Sprintf("INSERT INTO %s (embedding) VALUES ($1)", *tablename)
		if _, err := conn.Exec(ctx, insertSQL, vector); err != nil {
			log.Fatalf("Error inserting vector: %v", err)
		}

		inserted++
		// Print progress every 1000 vectors.
		if inserted%1000 == 0 {
			fmt.Printf("%d vectors inserted...\n", inserted)
		}
	}

	fmt.Printf("Finished inserting %d vectors into table '%s'.\n", inserted, *tablename)
}
