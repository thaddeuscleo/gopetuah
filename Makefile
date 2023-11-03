run:
	@go run cmd/proxy/main.go

test-server-1:
	@go run cmd/test_server/main.go 8000

test-server-2:
	@go run cmd/test_server/main.go 9000