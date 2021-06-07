.PHONY: build
build: deps
	go build ./cmd/cvimg

.PHONY: test
test:
	go test ./cmd/cvimg
