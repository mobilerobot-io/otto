build:
	go build -v

all: 
	go build -v
	make -C echo
	make -C static

test:
	go test -v

run: all
	./otto echo/echo.so static/static.go


