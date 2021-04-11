PKG=github.com/gopherd/doge/build
BRANCH=$(shell git symbolic-ref --short HEAD)
HASH=$(shell git rev-parse HEAD)
DATE=$(shell date "+%Y/%m/%d")
TIME=$(shell date "+%H:%M:%S")
GOBUILD=go build -ldflags "-X ${PKG}.branch=${BRANCH} -X ${PKG}.hash=${HASH} -X ${PKG}.date=${DATE} -X ${PKG}.time=${TIME}"
BUILD = ./build

all: gopherd

init:
	@mkdir -p ${BUILD}
	@mkdir -p ${BUILD}/

gopherd: init
	@echo "Building gopherd"
	@${GOBUILD} -o ${BUILD}/gopherd ./cmd/gopherd/
