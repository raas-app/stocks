# Variables
MIGRATE=migrate
MIGRATION_DIR=migrations
UP_DIR=$(MIGRATION_DIR)/up
DOWN_DIR=$(MIGRATION_DIR)/down
DATABASE_URL=mysql://root@tcp(localhost:3306)/stocks

.PHONY: create
create: # Create Migration (pass the name of the migration as the argument)
	$(MIGRATE) create -ext sql -dir $(MIGRATION_DIR) $(NAME)
	mv $(MIGRATION_DIR)/*.up.sql $(UP_DIR)/
	mv $(MIGRATION_DIR)/*.down.sql $(DOWN_DIR)/

# Apply Migrations (apply all pending migrations from 'up' directory)
up:
	@echo "Running migrations up..."
	$(MIGRATE) -path $(UP_DIR) -database "$(DATABASE_URL)" up

# Apply Specific Migration (apply a specific migration by ID from 'up' directory)
up-to:
	@echo "Running migrations up to $(ID)..."
	$(MIGRATE) -path $(UP_DIR) -database "$(DATABASE_URL)" goto $(ID)

# Rollback Migrations (revert the last migration from 'down' directory)
down:
	@echo "Rolling back the last migration..."
	$(MIGRATE) -path $(DOWN_DIR) -database "$(DATABASE_URL)" down

# Rollback Specific Number of Migrations from 'down' directory
down-to:
	@echo "Rolling back $(N) migrations..."
	$(MIGRATE) -path $(DOWN_DIR) -database "$(DATABASE_URL)" down $(N)

# Show Current Migration Version
version:
	@echo "Showing current migration version..."
	$(MIGRATE) -path $(MIGRATION_DIR) -database "$(DATABASE_URL)" version

# Make sure the migrate binary is installed (Optional)
install:
	@echo "Installing migrate tool..."
	go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Run all the above migrations (example)
all: up version
