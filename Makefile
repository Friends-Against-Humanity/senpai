dev:
	@go run cmd/main.go

test:
	@go test ./...

test-coverage:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out
