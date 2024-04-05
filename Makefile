.PHONY: up down start stop restart

up:
	docker-compose -f deployments/docker-compose.yml --project-name cars up

down:
	docker-compose -f deployments/docker-compose.yml --project-name cars down

start:
	docker-compose -f deployments/docker-compose.yml --project-name cars start

stop:
	docker-compose -f deployments/docker-compose.yml --project-name cars stop

restart:
	docker-compose -f deployments/docker-compose.yml --project-name cars restart
