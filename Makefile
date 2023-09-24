build:
	docker compose -f docker-compose.dev.yml build

dev:
	docker compose -f docker-compose.dev.yml up -d

down:
	docker compose -f docker-compose.dev.yml down

test-up:
	docker compose -f docker-compose.test.yml up --build -d

test-run:
	./docker/migrate_test.sh

test-down:
	docker compose -f docker-compose.test.yml down --volumes

