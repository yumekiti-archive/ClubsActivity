dc := docker compose -f ./docker-compose.yml

up:
	$(dc) up -d

down:
	$(dc) down

restart:
	$(dc) restart

reup:
	@make down
	@make up

rm:
	$(dc) down --rmi all

logs:
	$(dc) logs -f

web:
	$(dc) exec web /bin/sh

api:
	$(dc) exec api /bin/sh

format:
	$(dc) exec web yarn format

.PHONY:	up down restart reup rm logs web api format