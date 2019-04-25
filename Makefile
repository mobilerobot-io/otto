# plugins ::= $(wildcard $(dir $(wildcard */*.so)))

all: build $(plugins) pi build

status:
	@echo "All good with OttO"

build:
	go build -v

pi:
	env GOOS=linux GOARCH=arm GOARM=5 go build -v
	$(MAKE) -C plugins pi

run:
	make run -v main.go

.PHONY: $(plugins) build plugins all
