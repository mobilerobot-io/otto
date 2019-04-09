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
	make -C otto run

.PHONY: $(plugins) build
