plugins := $(wildcard $(dir $(wildcard */*.so)))

build: otto plugins

otto: *.go
	go build -v

plugins:
	$(MAKE) -C plugins

pi:
	env GOOS=linux GOARCH=arm GOARM=5 go build -v
	$(MAKE) -C plugins pi

run: otto
	./otto ${OFLAGS}

status:
	@echo "All good with OttO"

clean:
	go clean
	rm -rf *~ otto *.so

.PHONY: $(plugins) build plugins all
