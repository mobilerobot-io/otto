build:
	go build -v

test:
	go test -v

run: build
	go run main.go 

