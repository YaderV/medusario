build:
	docker compose -f docker-compose.dev.yml build

dev:
	docker compose -f docker-compose.dev.yml up -d

down:
	docker compose -f docker-compose.dev.yml down

test:
	docker compose -f docker-compose.test.yml up --build

test-quick:
	docker compose -f docker-compose.test.yml up

test-cleanup:
	docker compose -f docker-compose.test.yml down --volumes --remove-orphans --rmi all

