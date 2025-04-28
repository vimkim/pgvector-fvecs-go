example:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=vectbl -fvecs ../vector-gen-go/data/vector_256dim_3row_seed0.fvecs

load-128-75000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_128_75000 -fvecs ../vector-gen-go/data/vector_128dim_75000row_seed0.fvecs

load-128-150000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_128_150000 -fvecs ../vector-gen-go/data/vector_128dim_150000row_seed0.fvecs

load-128-300000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_128_300000 -fvecs ../vector-gen-go/data/vector_128dim_300000row_seed0.fvecs

###############################################################################
# load
###############################################################################

load-all: load-256-all load-768-all load-1536-all

load-256-all: load-256-75000 load-256-150000 load-256-300000

load-768-all: load-768-75000 load-768-150000 load-768-300000

load-1536-all: load-1536-75000 load-1536-150000 load-1536-300000

load-256-75000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_256_75000 -fvecs ../vector-gen-go/data/vector_256dim_75000row_seed0.fvecs

load-256-150000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_256_150000 -fvecs ../vector-gen-go/data/vector_256dim_150000row_seed0.fvecs

load-256-300000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_256_300000 -fvecs ../vector-gen-go/data/vector_256dim_300000row_seed0.fvecs

load-400-75000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_400_75000 -fvecs ../vector-gen-go/data/vector_400dim_75000row_seed0.fvecs

load-500-75000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_500_75000 -fvecs ../vector-gen-go/data/vector_500dim_75000row_seed0.fvecs

load-501-75000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_501_75000 -fvecs ../vector-gen-go/data/vector_501dim_75000row_seed0.fvecs

load-512-3:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_512_3 -fvecs ../vector-gen-go/data/vector_512dim_3row_seed0.fvecs

load-512-75000:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=tbl_512_75000 -fvecs ../vector-gen-go/data/vector_512dim_75000row_seed0.fvecs

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

load-nytimes-train:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=nytimes_256_angular_train -fvecs ../cubvec-61-perftest/nytimes-256-angular-train.fvecs

load-nytimes-test:
    go run load_fvecs.go --dbname=vimkimdb -user=vimkim -tablename=nytimes_256_angular_test -fvecs ../cubvec-61-perftest/nytimes-256-angular-test.fvecs
