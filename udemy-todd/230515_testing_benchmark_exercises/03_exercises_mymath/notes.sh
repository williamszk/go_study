# install some command line tools for making better Go code
go install golang.org/x/tools/cmd/godoc@latest
go install golang.org/x/lint/golint@latest

# go test ===================================================
# to run all tests
go test ./...
# we go enter into specific dirs to run the specific packages that we want

# check the coverage of tests
go test -cover ./...

# godocs ===================================================
# to spin up the godocs web page
$(go env GOPATH)/bin/godoc -http=:8080
# if the gopath bin is already in the PATH then we can just run the line below:
godoc -http=:8080

# golint ===================================================
# this will call out anything that is not idiomatic
$(go env GOPATH)/bin/golint ./...
golint ./...

# go bench ===================================================
# 
cd mymath
go test -bench .
cd -