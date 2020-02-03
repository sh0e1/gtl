test:
	go test -v -race ./...

lint:
	@if ! type golangci-lint; then \
		go get -u github.com/golangci/golangci-lint; \
	fi
	golangci-lint run ./... -E misspell
