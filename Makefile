


all: build 
.PHONY: all 

build:
	go build $(GOFLAGS) -o ./bin/aquarium  main.go
.PHONY: build