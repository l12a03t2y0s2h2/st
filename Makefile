.SILENT:
.DEFAULT_GOAL=lint

BINARY_NAME=st
BINARY_DBG_NAME=st_dbg
BINARY_PATH=bin

lint:
	golangci-lint run

# build: without debugging info -ldflags "-s -w" (go tool link)
build:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $(BINARY_PATH)/$(BINARY_NAME) .

# build: with debugging info
build_dbg:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_PATH)/$(BINARY_DBG_NAME) .

build_all: build build_dbg

size: build_all
	du -h $(BINARY_PATH)/*

copy: build
	sudo cp $(BINARY_PATH)/$(BINARY_NAME) /usr/local/bin
