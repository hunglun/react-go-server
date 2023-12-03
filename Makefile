all:
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags="-extldflags=-static" -o bin/arm_webserver
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-extldflags=-static" -o bin/arm64_webserver
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-extldflags=-static" -o bin/amd64_webserver

x86:
	go build -ldflags="-extldflags=-static" -o bin/x86_webserver
arm:
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -ldflags="-extldflags=-static" -o bin/arm_webserver

arm64:
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-extldflags=-static" -o bin/arm64_webserver

amd64:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-extldflags=-static" -o bin/amd64_webserver

.PHONY: run
run:
	go run main.go