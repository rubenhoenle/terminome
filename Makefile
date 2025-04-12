build:
	go build -o ./bin/terminome ./cmd/terminome/main.go 

test:
	go test ./...

lint:
	@echo ">> Linting with golangci-lint"
	golangci-lint run

fmt:
	gofmt -w .
