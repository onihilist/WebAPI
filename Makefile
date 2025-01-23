build:
	@docker compose up --build

clean-restart:
	@docker compose down mariadb -v
	@docker compose down app -v
	@docker image prune -a -f
	@docker network prune -f
	@docker volume prune -a -f
	@docker compose up mariadb -d
	@docker compose up app

up:
	@docker compose up -d

down:
	@docker compose down

restart:
	@docker compose restart

downwipe:
	@docker compose down -v

pgshell:
	@docker compose exec pgdb bash

psql:
	@docker compose exec pgdb psql -h pgdb -d appdb -U postgres

myshell:
	@docker compose exec mariadb bash

mysql:
	@docker compose exec mariadb mysql --host mariadb --database appdb --user appuser --password
