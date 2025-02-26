
include ./app/pkg/config/.env
export $(shell sed 's/=.*//' ./app/pkg/config/.env)

setup-env:
	@> ./app/pkg/config/.env
	@for file in ./docker.conf/*.env; do \
		cat "$$file" >> ./app/pkg/config/.env; \
		echo "" >> ./app/pkg/config/.env; \
	done

build:
	@docker compose up --build

up:
	@docker compose --env-file ./app/pkg/config/.env up -d

down:
	@docker compose down -v

clean:
	@docker image prune -a -f
	@docker network prune -f
	@docker volume prune -a -f

restart:
	@make down
	@make clean
	@make up

restartapp:
	@docker stop $(APP_CONTAINER_NAME)
	@docker remove $(APP_CONTAINER_NAME)
	@echo "Removing image $(APP_IMAGE_NAME)..."
	docker rmi $(APP_IMAGE_NAME) || true
	@docker compose up app -d

help:
	@echo "Available commands:"
	@echo ""
	@echo "  make setup-env"
	@echo "    - Combine all .env files from ./docker.conf/ into a single .env file at ./app/pkg/config/.env."
	@echo ""
	@echo "  make build"
	@echo "    - Build and start all containers defined in docker-compose.yml."
	@echo ""
	@echo "  make up"
	@echo "    - Start all containers in detached mode using the combined .env file."
	@echo ""
	@echo "  make down"
	@echo "    - Stop and remove all containers, networks, and volumes."
	@echo ""
	@echo "  make clean"
	@echo "    - Clean up unused Docker resources (images, networks, volumes)."
	@echo ""
	@echo "  make restart"
	@echo "    - Restart the entire setup: stop, clean, and start all containers."
	@echo ""
	@echo "  make restartapp"
	@echo "    - Restart only the application container: stop, remove, and restart the app container."
	@echo ""
	@echo "  make help"
	@echo "    - Display this help message."
	@echo ""