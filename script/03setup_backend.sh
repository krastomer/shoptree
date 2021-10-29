docker build -f backend/Dockerfile -t st-backend ./backend
docker run -p 8080:8080 --rm --network=host st-backend