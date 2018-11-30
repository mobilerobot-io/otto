build:
	go build -v

all: 
	go build -v
	make -C echo

test:
	go test -v

run: all
	./otto echo/echo.so 


