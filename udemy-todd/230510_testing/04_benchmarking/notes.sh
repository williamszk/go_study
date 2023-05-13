
go run main.go

go test ./...

# test just the sayings dir
go test ./sayings

# open godoc
godoc -http=:8080

# run the benchmarks
cd sayings
go test -bench .
cd -

# if we want to specify the benchmark to run
cd sayings
go test -bench Greet
cd -

# about coverage
go test -cover
# we can run this inside the specific package
# like inside the "sayings" directory

# or we can run in all directories down the current one
go test -cover ./...
# if you run this from root we'll test the whole project 

go test ./... -coverprofile c.out
go tool cover -html=c.out # this didn't work...

