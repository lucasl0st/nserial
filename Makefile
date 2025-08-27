.PHONY: fmt
fmt:
	go tool mvdan.cc/gofumpt -w .

.PHONY: generate
generate:
	go generate ./...

.PHONY: lint
lint:
	go tool github.com/golangci/golangci-lint/v2/cmd/golangci-lint run --timeout 5m

.PHONY: unit-test
unit-test:
	go test -tags unit -v -count=1 ./...
