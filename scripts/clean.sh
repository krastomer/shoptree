#!/bin/sh

docker-compose down
docker image prune -a
docker volume prune