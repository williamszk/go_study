
# to run the main 
go run main.go

# to run the tests
go test ./...

# to run the docs
godoc -http=:8080

# to run the benchmarks
cd word
go test -bench .
cd -

# check coverage of tests
go test -cover ./...


