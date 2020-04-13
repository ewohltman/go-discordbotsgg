.PHONY: lint test

lint:
	golangci-lint run ./...

test:
	go test -v -race -coverprofile=coverage.out ./...

bench:
	go test -v -run=X -bench=. -race -coverprofile=coverage.out ./...
