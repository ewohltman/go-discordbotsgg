.PHONY: lint test

lint:
	golangci-lint run ./...

test:
	go test -v -bench=. -race -coverprofile=coverage.out ./...
