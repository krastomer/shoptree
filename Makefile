all: dev

dev:	
	docker-compose -f docker-compose.dev.yml build --parallel
	docker-compose -f docker-compose.dev.yml up

production:
	docker-compose up --build

clean:
	docker-compose -f docker-compose.dev.yml rm
	docker-compose rm
	docker image prune -a