DB_URL_LOCAL = postgres://smap_user:smap_password@localhost:5432/smap?sslmode=disable
DB_URL_DOCKER = postgres://smap_user:smap_password@postgres:5432/smap?sslmode=disable
MIGRATIONS_DIR = ./migrations

.PHONY: migrate-up migrate-down migrate-create migrate-version migrate-force

migrate-up:
	@echo "Applying migration"
	@migrate -path $(MIGRATIONS_DIR) -database $(DB_URL_LOCAL) up

migrate-down:
	@echo "Rolling back migration"
	@migrate -path $(MIGRATIONS_DIR) -database $(DB_URL_LOCAL) down 1

migrate-create:
	@if [ -z "$(name)" ]; then
	    echo "Error: migration name is required. Usage: make migrate-create name=%name%"
		exit 1
	fi
	@migrate create -ext sqal -dir $(MIGRATIONS_DIR) -seq $(name)

migrate-up-docker:
	@echo "Applying migrations inside Docker network..."
	@docker-compose run --rm user-service migrate -path /app/migrations -database $(DB_URL_DOCKER) up

migrate-down-docker:
	@echo "Rolling back migrations inside Docker"
	@docker-compose run --rm user-service migrate -path /app/migrations -database $(DB_URL_DOCKER) down 1

migrate-version:
	@migrate -path $(MIGRATIONS_DIR) -database $(DB_URL_LOCAL) version

migrate-force:
	@if [ -z "$(version)" ]; then \
		echo "Error: version is required. Usage: make migrate-force version=1"; \
		exit 1; \
	fi
	@migrate -path $(MIGRATIONS_DIR) -database $(DB_URL_LOCAL) force $(version)

help:
	@echo "Available migration commands:"
	@echo "  migrate-up                        - start migration"
	@echo "  migrate-down                      - rollback migration one step down"
	@echo "  migrate-create                    - create new migration"
	@echo "  make migrate-version              - Show current version"
	@echo "  make migrate-force version=N      - Force set migration version"