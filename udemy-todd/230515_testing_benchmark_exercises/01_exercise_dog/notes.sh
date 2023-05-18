
# [ ] Tests
# [ ] Examples
# [ ] Benchmarks
# [ ] Coverage

# Can you BET on your code?

# to run the main 
go run main.go

# to run the tests
go test ./...

# to run the docs
godoc -http=:8080

# to run the benchmarks
cd dog
go test -bench .
cd -

# check coverage of tests
go test -cover ./...

