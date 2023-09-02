build:
	@go build -o bin/goFactorial

run: build
	@./bin/goFactorial

test:
	@go test -v ./...