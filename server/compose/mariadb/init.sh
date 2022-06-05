#!/bin/sh
docker volume create mariadb_data
docker volume create mariadb_logs
docker run -v mariadb_data:/tmp busybox chown -R 27:27 /tmp
docker run -v mariadb_logs:/tmp busybox chown -R 27:27 /tmp
