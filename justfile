example:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=vectbl -fvecs ../vector-gen-go/data/vector_256dim_3row_seed0.fvecs

load-128-75000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_128_75000 -fvecs ../vector-gen-go/data/vector_128dim_75000row_seed0.fvecs

load-128-150000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_128_150000 -fvecs ../vector-gen-go/data/vector_128dim_150000row_seed0.fvecs

load-128-300000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_128_300000 -fvecs ../vector-gen-go/data/vector_128dim_300000row_seed0.fvecs

load-256-75000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_256_75000 -fvecs ../vector-gen-go/data/vector_256dim_75000row_seed0.fvecs

load-256-150000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_256_150000 -fvecs ../vector-gen-go/data/vector_256dim_150000row_seed0.fvecs

load-256-300000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_256_300000 -fvecs ../vector-gen-go/data/vector_256dim_300000row_seed0.fvecs

load-768-all: load-768-75000 load-768-150000 load-768-300000

load-768-75000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_768_75000 -fvecs ../vector-gen-go/data/vector_768dim_75000row_seed0.fvecs

load-768-150000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_768_150000 -fvecs ../vector-gen-go/data/vector_768dim_150000row_seed0.fvecs

load-768-300000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_768_300000 -fvecs ../vector-gen-go/data/vector_768dim_300000row_seed0.fvecs

load-1536-75000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_1536_75000 -fvecs ../vector-gen-go/data/vector_1536dim_75000row_seed0.fvecs

load-1536-150000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_1536_150000 -fvecs ../vector-gen-go/data/vector_1536dim_150000row_seed0.fvecs

load-1536-300000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_1536_300000 -fvecs ../vector-gen-go/data/vector_1536dim_300000row_seed0.fvecs
