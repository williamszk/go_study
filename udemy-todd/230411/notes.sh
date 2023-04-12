# https://www.udemy.com/course/learn-how-to-code/learn/lecture/11921970#overview

# About the concepts of Go Workspace

cd 01
mkdir goworkspace
cd goworkspace

mkdir bin pkg src

pwd
# /root/go_study/udemy-todd/230411/01/goworkspace

# Check Go environment variables
# This is the original go environment
go env
# GO111MODULE=""
# GOARCH="amd64"
# GOBIN=""
# GOCACHE="/root/.cache/go-build"
# GOENV="/root/.config/go/env"
# GOEXE=""
# GOEXPERIMENT=""
# GOFLAGS=""
# GOHOSTARCH="amd64"
# GOHOSTOS="linux"
# GOINSECURE=""
# GOMODCACHE="/root/go/pkg/mod"
# GONOPROXY=""
# GONOSUMDB=""
# GOOS="linux"
# GOPATH="/root/go"
# GOPRIVATE=""
# GOPROXY="https://proxy.golang.org,direct"
# GOROOT="/usr/local/go"
# GOSUMDB="sum.golang.org"
# GOTMPDIR=""
# GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
# GOVCS=""
# GOVERSION="go1.19.3"
# GCCGO="gccgo"
# GOAMD64="v1"
# AR="ar"
# CC="gcc"
# CXX="g++"
# CGO_ENABLED="1"
# GOMOD="/dev/null"
# GOWORK=""
# CGO_CFLAGS="-g -O2"
# CGO_CPPFLAGS=""
# CGO_CXXFLAGS="-g -O2"
# CGO_FFLAGS="-g -O2"
# CGO_LDFLAGS="-g -O2"
# PKG_CONFIG="pkg-config"
# GOGCCFLAGS="-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build4281111270=/tmp/go-build -gno-record-gcc-switches"

# This is an important one
# GOPATH="/root/go"
GOPATH="/root/go"

export GOPATH="/root/go_study/udemy-todd/230411/01/goworkspace"
echo $GOPATH
# export GOPATH=$HOME/go

go env
# root@554703e9de47:~/go_study/udemy-todd/230411/01/goworkspace# go env
# GOARCH="amd64"
# ...
# GOOS="linux"
# GOPATH="/root/go_study/udemy-todd/230411/01/goworkspace" <---

# If I close the terminal and open it again we get:
# GOPATH="/root/go"



