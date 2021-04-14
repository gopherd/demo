PKG=github.com/gopherd/doge/build
BRANCH=$(shell git symbolic-ref --short HEAD)
HASH=$(shell git rev-parse HEAD)
DATE=$(shell date "+%Y/%m/%d")
TIME=$(shell date "+%H:%M:%S")
GOBUILD=go build -ldflags "-X ${PKG}.branch=${BRANCH} -X ${PKG}.hash=${HASH} -X ${PKG}.date=${DATE} -X ${PKG}.time=${TIME}"
BUILDDIR = ./build

all: ram rem

init:
	@mkdir -p ${BUILDDIR}
	@mkdir -p ${BUILDDIR}/

ram: init
	@echo "Building ram"
	@${GOBUILD} -o ${BUILDDIR}/ram ./cmd/ram/

rem: init
	@echo "Building rem"
	@${GOBUILD} -o ${BUILDDIR}/rem ./cmd/rem/
