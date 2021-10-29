#!bin/bash

docker exec -i shoptree-database sh -c 'exec mysql -uroot -ppassword' < ./database/shoptree.sql