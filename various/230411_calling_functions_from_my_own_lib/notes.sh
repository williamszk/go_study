
go mod init calling_my_own_lib

go clean -modcache

# -u is for updating
go get github.com/williamszk/a_lib_in_go
go get -u github.com/williamszk/a_lib_in_go
go get -u github.com/williamszk/a_lib_in_go@v1.0.1
# check later inside the go env GOPATH
# if your lib is installed there

go build
./calling_my_own_lib

# It didn't work...

# root@554703e9de47:~/go_study/various/230411_calling_functions_from_my_own_lib# go get github.com/williamszk/a_lib_in_go
# go: github.com/williamszk/a_lib_in_go@v0.0.0-20230411174757-8df8a0f1ceef: parsing go.mod:
#         module declares its path as: a_lib_in_go
#                 but was required as: github.com/williamszk/a_lib_in_go

# I'll try again another day.