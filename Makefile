build:
	@docker compose up --build

clean:
	@docker image prune -a -f
	@docker network prune -f
	@docker volume prune -a -f

up:
	@docker compose up -d

down:
	@docker compose down -v

delete:
	@docker rm mariadb
	@docker rm app

restart:
	@docker compose restart
