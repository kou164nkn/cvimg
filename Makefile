.PHONY: build
build:
	go build ./cmd/cvimg

.PHONY: test
test:
	go test -v -race -cover
