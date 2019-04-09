# plugins ::= $(wildcard $(dir $(wildcard */*.so)))

plugins =  echo wally static store dork

all: $(plugins) build

$(plugins):
	${MAKE} -C $@ 

status:
	@echo "All good with OttO"

build:
	go build

run:
	make run -v main.go

.PHONY: $(plugins) build
