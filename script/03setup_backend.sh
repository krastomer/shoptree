#!/bin/bash

docker build -f backend/Dockerfile -t shoptree-backend ./backend
docker run -p 8080:8080 --restart=always shoptree-backend