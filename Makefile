.PHONY: lint
lint:
	gometalinter.v3 ./...

.PHONY: test
test:
	go test -v -cover ./...
