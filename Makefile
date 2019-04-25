# plugins ::= $(wildcard $(dir $(wildcard */*.so)))

build:
	go build -v
	$(MAKE) -C plugins

pi:
	env GOOS=linux GOARCH=arm GOARM=5 go build -v
	$(MAKE) -C plugins pi

run:
	make run -v main.go

status:
	@echo "All good with OttO"

clean:
	go clean
	rm -rm *~ otto *.so


.PHONY: $(plugins) build plugins all
