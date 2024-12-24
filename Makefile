# Vars
# replace with the name and other vars of your application
IMAGE_NAME ?= v3/go-template
IMAGE_TAG ?= 0.1.0
BIN_NAME ?= main
BIN_PATH ?= ./bin
CMD_PATH ?= ./cmd

# Tasks
build: clean deps doc
	go build -o $(BIN_PATH)/$(BIN_NAME) $(CMD_PATH)/$(BIN_NAME).go

clean:
	rm -rf $(BIN_PATH)

deps:
	go mod tidy

dev:
	go run $(CMD_PATH)/$(BIN_NAME).go

doc:
	swag init -g swagger.go -d ./internal -o ./api/docs --parseDependency --parseInternal 

test:
	go test -v ./...

build-image:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

docker-up: docker-down
	docker compose up

docker-down:
	docker compose down -v --remove-orphans

start: build
	$(BIN_PATH)/$(BIN_NAME)
