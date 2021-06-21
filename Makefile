PKG=github.com/gopherd/doge/build
BRANCH=$(shell git symbolic-ref --short HEAD)
HASH=$(shell git rev-parse HEAD)
DATE=$(shell date "+%Y/%m/%d")
TIME=$(shell date "+%H:%M:%S")
GOBUILD=go build -trimpath -ldflags "-X ${PKG}.branch=${BRANCH} -X ${PKG}.hash=${HASH} -X ${PKG}.date=${DATE} -X ${PKG}.time=${TIME}"
TARGET_DIR=./target

define build_target
	@mkdir -p ${TARGET_DIR}
	@echo "Building ${TARGET_DIR}/$(1) ..."
	@${GOBUILD} -o ${TARGET_DIR}/$(1) ./cmd/$(1)/
endef

all: ram rem

ram:
	$(call build_target,ram)

rem:
	$(call build_target,rem)
