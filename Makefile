.SILENT:
.DEFAULT_GOAL=lint

BINARY_NAME=st
BINARY_DBG_NAME=st_dbg
BINARY_PATH=bin

lint:
	goimports -l -w .
	golangci-lint run

# build: without debugging info -ldflags "-s -w" (go tool link)
binary:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $(BINARY_PATH)/$(BINARY_NAME) .

# build: with debugging info
build_dbg:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_PATH)/$(BINARY_DBG_NAME) .

build_all: binary build_dbg

size: build_all
	du -h $(BINARY_PATH)/*

copy: binary
	sudo cp $(BINARY_PATH)/$(BINARY_NAME) /usr/local/bin
