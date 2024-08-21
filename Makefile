server:
	@echo "Starting server..."
	@go run main.go

postgres:
	@echo "Starting postgres..."
	@docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres:16-alpine

createdb:
	@echo "Creating database..."
	@docker exec -it postgres createdb --username=postgres --owner=root project8

dropdb:
	@echo "Dropping database..."
	@docker exec -it postgres dropdb project8

sqlc:
	@echo "Generating sqlc..."
	@sqlc generate

migration:
	@echo "Creating migration..."
	@migrate create -ext sql -dir db/migration $(name)

.PHONY: server postgres createdb dropdb sqlc