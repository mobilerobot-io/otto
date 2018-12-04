build:
	go build -v

all: 
	go build -v
	make -C echo
	make -C static
	make -C store
	make -C dork
	make -C wally

test:
	go test -v

run: all
	./otto dork/dork.so

dork: 
	make -C dork 
	go build -v && ./otto dork/dork.so 

wally: 
	make -C wally 
	go build -v && ./otto wally/wally.so 

.PHONY: dork wally echo static