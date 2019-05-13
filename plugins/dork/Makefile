src = dork.so
plugin = dork.so

build: $(plugin)

$(plugin): $(src)
	go build -v -buildmode=plugin

pi:
	env GOOS=linux GOARCH=arm GOARM=5 go build -v -buildmode=plugin

.PHONEY: all build pi $(plugin)
