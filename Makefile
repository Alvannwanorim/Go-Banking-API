build: 
	@go build -o bin/Go-Banking-API

run: build
	@./bin/Go-Banking-API

test:
	@go test -v ./...