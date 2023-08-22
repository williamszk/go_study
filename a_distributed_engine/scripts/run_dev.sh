#! /bin/bash

docker build -t my_app -f Dockerfile.dev .
docker run --rm -it my_app "./myapp worker"
