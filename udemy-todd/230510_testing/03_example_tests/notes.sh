
go run main.go

go test

# ./... means that we want to run all the dirs down the current one
# it is like the recursive option
go test ./...

# generate documentation
go get golang.org/x/tools/cmd/godoc
go install golang.org/x/tools/cmd/godoc@latest
echo $GOPATH
$GOPATH/bin/godoc -http=:6060
godoc -http=:8080


$(go env GOPATH)/bin/godoc -http=:8080

# experimenting with golint ===================================================

go install golang.org/x/tools/cmd/godoc@latest
go install golang.org/x/lint/golint@latest
$(go env GOPATH)/bin/golint ./...

# use go vet, why this exists?
go vet ./...

