# Commands related to the database
db-start:
	docker-compose up -d

db-stop:
	docker-compose stop

db-kill:
	docker-compose down -v

# Update database schema after an update
generate-schema:
	go generate ./ent

# Start server
go-server:
	go run api/main.go

go-test:
	go test ./...

go-format: # Run go format to format files
	go fmt ./...