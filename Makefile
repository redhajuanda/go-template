include .env
export

psql:
	@PGPASSWORD=${DB_PASSWORD} psql -h ${DB_HOST} -U ${DB_USERNAME}

.PHONY: migrate
migrate:
	@sql-migrate up

.PHONY: migrate-redo
migrate:
	@sql-migrate redo

.PHONY: migrate-fresh
migrate-fresh:
	@echo "Resetting database..."
	@PGPASSWORD="$$DB_PASSWORD" psql -h "$$DB_HOST" -U "$$DB_USERNAME" -d "$$DB_NAME" -c 'drop schema public cascade; create schema public;'
	@sql-migrate up

.PHONY: migrate-new
migrate-new: ## create a new database migration
	@read -p "Enter the name of the new migration: " name; \
	sql-migrate new $${name}

run:
	@go run main.go