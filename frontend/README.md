# Docker command

`docker build -f Dockerfile.dev -t treeshop:dev`

`docker run -it --rm -v ${PWD}:/app -v /app/node_modules -p 3000:3000 -e CHOKIDAR_USEPOLLING=true treeshop:dev`
