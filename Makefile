build: clean deps
	go build -o ./bin/main ./cmd/main.go

clean:
	rm -rf ./bin

deps:
	go mod tidy
	
dev:
	go run ./cmd/main.go

docker-up: docker-down
	docker compose up
	
docker-down:
	docker compose down -v --remove-orphans

start: build
	./bin/main