


all: build init
.PHONY: all 

init:
	go run tools/init_server/main.go
.PHONY: init

build:
	go build $(GOFLAGS) -o ./bin/aquarium  main.go
.PHONY: build