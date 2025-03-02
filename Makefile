.PHONY: run build down clean

clean:
	go fmt
	swag fmt
	swag init
	clear

dev-run: clean
	go run main.go

run:
	docker-compose up --build -d

down:
	docker-compose down

test:
	go test -v ./...

