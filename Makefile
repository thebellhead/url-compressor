.PHONY: test up down

test:
	go fmt ./...
	go mod tidy -v
	go clean -testcache
	go test -v ./...

up:
	docker compose -f ./docker-compose.yml rm && \
	docker compose -f ./docker-compose.yml build --no-cache && \
	docker compose -f ./docker-compose.yml up

down:
	docker-compose -f ./docker-compose.yml down
