# distributed engine with go = dengo
go mod init dengo
touch main.go

docker build -t my_app -f Dockerfile.dev .
docker run --rm -it my_app




