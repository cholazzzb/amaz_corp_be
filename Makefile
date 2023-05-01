dev:
	@go run main.go

test:
	@go test ./...

clean-test-cache:
	@go clean -testcache
