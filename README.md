# Roomie-go

This project aims to create an app for finding and contacting house/apartment hosts who are looking for roommates.
It is currently in development, following a microservices architecture, solely for educational purposes.

## Requirements
- Docker compose

## Deveploment
- Run `make build` to build all the project image
- Run `make dev` to start all the containers
- Run `make stop` to stop all the containers

## Testing
- Run `make test-up` to set the testing enviroment for unit and integration tests
- Run `make test-run` to run all the tests
- Run `make test-down` to remove the testing containers and volumes

## Docs
### Swagger
- <project-url>/v1/swagger.yaml

### Redoc
- <project-url>/v1/docs
