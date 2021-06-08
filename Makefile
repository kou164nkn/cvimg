.PHONY: build
build:
	go build ./cmd/cvimg

.PHONY: test
test:
	go test ./cmd/cvimg
