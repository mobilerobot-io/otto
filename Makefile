# plugins ::= $(wildcard $(dir $(wildcard */*.so)))

plugins =  echo wally static store dork

all: build $(plugins) build

$(plugins): 
	${MAKE} -C $@ 

status:
	@echo "All good with OttO"

build:
	go build -v

pi:
	env GOOS=linux GOARCH=arm GOARM=5 go build -v

run:
	make run -v main.go

.PHONY: $(plugins) build
