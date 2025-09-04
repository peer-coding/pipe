.PHONY: test
test: 
	go test ./... -v

.PHONY: build
build: 
	go build -o bin/app -trimpath -ldflags="-s -w" ./cmd/example/main.go
