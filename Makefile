USER_REGISTRATION_BINARY=userRegistrationApp
USER_BINARY=userApp
ADDRESS_BINARY=addressApp
USER_ADDRESS_AGG_BINARY=aggregatorApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose -f docker-compose.yaml -f docker-compose-mqtt.yaml up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_address-service build_user-registration build_user-service build_user-address-agg
#	@echo "Stopping docker images (if running...)"
#	docker-compose stop
#	docker-compose rm -f
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose stop
	docker-compose rm -f
	@echo "Done!"

## build_auth: builds the auth binary as a linux executable
build_user-registration:
	@echo "Building user-registration binary..."
	cd ./user-registration && env GOOS=linux CGO_ENABLED=0 go build -o ${USER_REGISTRATION_BINARY} ./cmd/api
	@echo "Done!"

build_user-service:
	@echo "Building user-service binary..."
	cd ./user-service && env GOOS=linux CGO_ENABLED=0 go build -o ${USER_BINARY} ./cmd/api
	@echo "Done!"

build_address-service:
	@echo "Building address-service binary..."
	cd ./address-service && env GOOS=linux CGO_ENABLED=0 go build -o ${ADDRESS_BINARY} ./cmd/api
	@echo "Done!"

build_user-address-agg:
	@echo "Building user-address-agg binary..."
	cd ./user-address-agg && env GOOS=linux CGO_ENABLED=0 go build -o ${USER_ADDRESS_AGG_BINARY} ./cmd/api
	@echo "Done!"
