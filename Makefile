run:
	@go run cmd/proxy/main.go

docker-build:
	@docker compose --profile build build
	@docker rmi -f $(docker images -f "dangling=true" -q)

docker-up:
	@docker compose --profile prod up -d

docker-down:
	@docker compose --profile prod down