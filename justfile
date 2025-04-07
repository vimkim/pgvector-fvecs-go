example:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=vectbl -fvecs ../vector-gen-go/data/vector_256dim_3row_seed0.fvecs

load-256-75000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_256_75000 -fvecs ../vector-gen-go/data/vector_256dim_75000row_seed0.fvecs

load-256-150000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_256_150000 -fvecs ../vector-gen-go/data/vector_256dim_150000row_seed0.fvecs

load-256-300000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_256_300000 -fvecs ../vector-gen-go/data/vector_256dim_300000row_seed0.fvecs
