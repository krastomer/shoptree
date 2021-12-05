#!/bin/sh

echo "When you want exit Ctrl + C 2 times"

mkdir -p ./data/images/products
mkdir -p ./data/images/payments

PATH_ADMIN="./backend/admin/app.env"
PATH_CUSTOMER="./backend/customer/app.env"
if ! test -f $PATH_ADMIN; then
    echo "APP_PORT=:8081
DB_HOST=mariadb
DB_USERNAME=root
DB_PASSWORD=password
DB_PORT=3306
DB_NAME=shoptree

BCRYPT_SIZE=8
JWT_SECRET=september

DIRECTORY_PRODUCT=./images/products
DIRECTORY_PAYMENT=./images/payments"> $PATH_ADMIN
fi

if ! test -f $PATH_CUSTOMER; then
    echo "APP_PORT=:8080
DB_HOST=mariadb
DB_USERNAME=root
DB_PASSWORD=password
DB_PORT=3306
DB_NAME=shoptree

BCRYPT_SIZE=8
JWT_SECRET=september

DIRECTORY_PRODUCT=./images/products
DIRECTORY_PAYMENT=./images/payments"> $PATH_CUSTOMER
fi

if ! type docker-compose; then
    echo docker-compose not found. Please install https://docs.docker.com/compose/install/.
    exit 1
fi

docker-compose build --parallel web-customer api-customer api-admin mariadb
docker-compose up web-customer api-customer api-admin mariadb
exit 0