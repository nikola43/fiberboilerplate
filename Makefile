build:
	@go build cmd/fiberboilerplate/main.go

run: build
	@./main

test:
	@go test -v ./...