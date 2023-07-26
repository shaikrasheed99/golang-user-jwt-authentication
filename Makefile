run:
	@go run .

run-dev:
	@compiledaemon --command="go run ."

test:
	@go test ./...

test-cov:
	@go test -tags 'test_coverage' -coverprofile=coverage.out ./...
	@go tool cover -func coverage.out

mocks:
	mockery --all --keeptree