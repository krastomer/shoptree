#!bin/sh

NAME_REPO="treeshop"
PORT=$1
NODE_VERSION=$(node -v)
DOCKER_VERSION=$(docker version --format '{{.Server.Version}}')

if [ ! $NODE_VERSION ]; then 
    echo "Didn't install node"
    exit 1
fi


if [ ! $DOCKER_VERSION ]; then
    echo "Didn't install docker"
    exit 1
fi

echo "node: $NODE_VERSION"
echo "docker: $DOCKER_VERSION"


echo "Find PORT"
if [ ! "$PORT" ]; then
    echo "Didn't find PORT. Let use default."
    PORT="80"
fi
echo "PORT is $PORT"

echo "Build React Production"
cd web
npm run-script build
echo "Done Builded..."

echo "Build Docker image: $NAME_REPO"
cd ..

docker build . -t $NAME_REPO

if [ ! "$(docker ps -q -f name=$NAME_REPO)" ]; then
    echo "Found container is exists."

    if [ "$(docker ps -aq -f status=running -f name=$NAME_REPO)" ]; then
        echo "Stop container."
        docker stop $NAME_REPO
    fi

    echo "Remove container."
    docker rm $NAME_REPO
fi

echo "Run container."
docker run --name=treeshop -dp 8080:8080 $NAME_REPO
echo "Done."
