test:
	go test -v ./cmd/... ./internal/...

build: 
	go build -o remote-numbers-fetcher ./cmd/
