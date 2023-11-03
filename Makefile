run:
	@go run cmd/proxy/main.go

docker-build:
	@docker compose --profile build build

docker-up:
	@docker compose --profile prod up -d

docker-down:
	@docker compose --profile prod down