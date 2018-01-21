.PHONY: emersyx_apis test

emersyx_apis:
	@echo "Running the tests with gofmt, go vet and golint..."
	@test -z $(shell gofmt -s -l emcomapi/*.go emircapi/*.go emtgapi/*.go emrtrapi/*.go)
	@go vet ./...
	@golint -set_exit_status $(shell go list ./...)
