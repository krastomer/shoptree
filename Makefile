all: dev

dev:	
	docker-compose build --parallel
	docker-compose up

clean:
	docker-compose -f docker-compose.dev.yml rm
	docker-compose rm
	docker image prune -a