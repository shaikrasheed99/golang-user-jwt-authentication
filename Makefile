run:
	@go run .

run-dev:
	@compiledaemon --command="go run ."

mocks:
	mockery --all --keeptree