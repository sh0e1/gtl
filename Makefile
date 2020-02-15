test:
	go test -v -race ./...

test/coverage:
	go test -v -race -cover ./...

lint:
	@if ! type golangci-lint; then \
		go get -u github.com/golangci/golangci-lint; \
	fi
	golangci-lint run ./... -E misspell
