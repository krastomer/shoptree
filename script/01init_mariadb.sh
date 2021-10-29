#!/bin/bash

database="shoptree-database"
dir="/app/shoptree/database"
password="password"
TZ="Asia/Bangkok"

docker run --name $database -v $dir:/var/lib/mysql -e MARIADB_ROOT_PASSWORD=$password -e TZ=$TZ --restart=true -dp 3306:3306 mariadb:10.5.12