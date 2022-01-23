# create a docker container for development of this project

sudo systemctl status docker

docker ps

# Got permission denied while trying to connect to the Docker daemon socket at 
# unix:///var/run/docker.sock: Get "http://%2Fvar%2Frun%2Fdocker.sock/v1.24/containers/json": dial unix /var/run/docker.sock: connect: permission denied

# https://www.digitalocean.com/community/questions/how-to-fix-docker-got-permission-denied-while-trying-to-connect-to-the-docker-daemon-socket

sudo groupadd docker

sudo usermod -aG docker $USER

sudo systemctl restart docker

# you may need to restart the machine

su -s william

docker run hello-world


# 

docker run --name go_study -it ubuntu /bin/bash

# inside the container ----------------------------
apt update
apt install git curl -y

# install go compiler
# download the go compiler
cd

curl -L https://go.dev/dl/go1.17.6.linux-amd64.tar.gz --output gobinary.tar.gz
# unzip the go compiler in the /usr/local/ directory
rm -rf /usr/local/go && tar -C /usr/local -xzf gobinary.tar.gz
# add the binary to the path so we can use in the terminal
export PATH=$PATH:/usr/local/go/bin

go version
# go version go1.17.6 linux/amd64

