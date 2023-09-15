build:
	docker compose -f dev.yml build

dev:
	docker compose -f dev.yml up -d

down:
	docker compose -f dev.yml down


