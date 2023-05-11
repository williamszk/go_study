# https://go.dev/blog/examples

# if godoc is not installed in your machine
go install golang.org/x/tools/cmd/godoc@latest

# to start the godoc page
$(go env GOPATH)/bin/godoc -http=:8081

go test ./... -v 