# Commands related to the database
start-db:
	docker-compose up -d

stop-db:
	docker-compose stop

kill-db:
	docker-compose down -v

# Commands to to update database schema
generate-schema:
	go generate ./ent

# Start
start-server:
	go run api/main.go

run-test:
	go test ./...