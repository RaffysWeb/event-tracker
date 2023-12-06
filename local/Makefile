CMD_PATH = ../cmd/
# Cassandra base path
CASSANDRA_URL = cassandra://localhost:9042/mykeyspace
MIGRATIONS_PATH = ../pkg/database/migrations/cql

# Binary base path
BINARY_BASE_PATH = ./bin/

# Define the binary name
EVENT_PRODUCER_BINARY = event-producer
EVENT_CONSUMER_BINARY = event-consumer

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_infra: starts all containers infra related resources
up_infra:
	@echo "Starting Docker images..."
	docker-compose up zookeeper kafka_broker cassandra influxdb telegraf -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_event_producer, build_event_consumer
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## up_build_infra: stops docker-compose (if running), builds and starts all infra related resources
up_build_infra:
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up zookeeper kafka_broker cassandra influxdb telegraf --build -d
	@echo "Docker images started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."own
	docker-compose down
	@echo "Done!"

## build_event_producer: builds the event producer binary as a linux executable
## Need o use Zig to compile with CGO_ENABLED=1
build_event_producer:
	@echo "Building event producer binary..."
	cd ${CMD_PATH}event-producer && \
		if [ $$(uname -m) = "aarch64" ]; then \
			env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC="zig cc -target x86_64-linux-musl" CXX="zig c++ -target x86_64-linux-musl" go build -tags musl -o ${BINARY_BASE_PATH}${EVENT_PRODUCER_BINARY}; \
		else \
			go build -o ${BINARY_BASE_PATH}${EVENT_PRODUCER_BINARY}; \
		fi
	@echo "Done!"

## build_event_consumer: builds the event producer binary as a linux executable
## Need o use Zig to compile with CGO_ENABLED=1
build_event_consumer:
	@echo "Building event consumer binary..."
	cd ${CMD_PATH}event-consumer && \
		if [ $$(uname -m) = "aarch64" ]; then \
			env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC="zig cc -target x86_64-linux-musl" CXX="zig c++ -target x86_64-linux-musl" go build -tags musl -o ${BINARY_BASE_PATH}${EVENT_CONSUMER_BINARY}; \
		else \
			go build -o ${BINARY_BASE_PATH}${EVENT_CONSUMER_BINARY}; \
		fi
	@echo "Done!"

# Migrate up
migrate_up:
	@echo "Migrating up..."
	journey --url ${CASSANDRA_URL} --path ${MIGRATIONS_PATH} migrate up
	@echo "Done!"

# Migrate down 1
migrate_down_1:
	@echo "Migrating down one..."
	journey --url ${CASSANDRA_URL} --path ${MIGRATIONS_PATH} migrate migrate -1
	@echo "Done!"

