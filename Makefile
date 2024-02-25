up:
	docker-compose -f docker/docker-compose.yml build --no-cache app
	docker-compose -f docker/docker-compose.yml up -d

down:
	docker-compose -f docker/docker-compose.yml down
