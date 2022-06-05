#!/bin/sh

docker-compose down
docker rm -f $(docker ps -a -q)
docker volume rm -f mariadb_data_mydb mariadb_logs_mydb
