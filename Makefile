build:
	@docker compose up --build

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
