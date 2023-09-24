#!/bin/bash

if [ -f .env ]; then
  export $(cat .env | xargs)
fi

CMD="docker compose -f docker-compose.test.yml --env-file ./.env run migrate-test \
    -path /migrations \
    -database postgres://\$DB_USER_TEST:\$DB_PASS_TEST@db-test:5432/\$DB_NAME_TEST?sslmode=disable \
    down -all"

eval $CMD

CMD="docker compose -f docker-compose.test.yml --env-file ./.env run migrate-test \
    -path /migrations \
    -database postgres://\$DB_USER_TEST:\$DB_PASS_TEST@db-test:5432/\$DB_NAME_TEST?sslmode=disable \
    up"

eval $CMD

docker compose -f docker-compose.test.yml exec test  go test ./... -v -count=1
