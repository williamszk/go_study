# https://go.dev/doc/gopath_code

mkdir bin src

export GOPATH=$PWD
echo $GOPATH
go env

go mod init go_workspace_experiment

touch main.go

go env GOPATH
echo $GOPATH

mkdir -p $GOPATH/src/github.com/user/hello
touch $GOPATH/src/github.com/user/hello/hello.go

cd $GOPATH/src/github.com/user/hello
go mod init github.com/user/hello
go install
cd -

$GOPATH/bin/hello
# Hello, world.

export PATH=$PATH:$(go env GOPATH)/bin
hello
# root@554703e9de47:~/go_study/site_go_dev/230411_go_path# hello
# Hello, world.



