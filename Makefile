BUILD = ./build

all: gopherd

init:
	@mkdir -p ${BUILD}
	@mkdir -p ${BUILD}/

gopherd: init
	go build -o ${BUILD}/gopherd ./cmd/gopherd/
