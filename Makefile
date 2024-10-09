.PHONY: migrate-up migrate-down migrate-force up up-build test down-volume

migrate-up:
	docker compose -f deployments/docker-compose.yaml --profile tools run migrate up

migrate-down:
	docker compose -f deployments/docker-compose.yaml --profile tools run migrate down

migrate-force:
	docker compose -f deployments/docker-compose.yaml --profile tools run migrate force $1

up:
	docker compose -f deployments/docker-compose.yaml up

up-build:
	docker compose -f deployments/docker-compose.yaml up --build

down-volume:
	docker compose -f deployments/docker-compose.yaml down --volumes

test:
	docker compose -f deployments/docker-compose.test.yaml up migrate && docker compose -f deployments/docker-compose.test.yaml up backend-test --build --remove-orphans && docker compose -f deployments/docker-compose.test.yaml down --volumes