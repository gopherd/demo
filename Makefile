PKG=github.com/gopherd/doge/build
BRANCH=$(shell git symbolic-ref --short HEAD)
HASH=$(shell git rev-parse HEAD)
DATE=$(shell date "+%Y/%m/%d")
TIME=$(shell date "+%H:%M:%S")
GOBUILD=go build -ldflags "-X ${PKG}.branch=${BRANCH} -X ${PKG}.hash=${HASH} -X ${PKG}.date=${DATE} -X ${PKG}.time=${TIME}"
BUILD = ./build

all: ram rem

init:
	@mkdir -p ${BUILD}
	@mkdir -p ${BUILD}/

ram: init
	@echo "Building ram"
	@${GOBUILD} -o ${BUILD}/ram ./cmd/ram/

rem: init
	@echo "Building rem"
	@${GOBUILD} -o ${BUILD}/rem ./cmd/rem/
