#!/bin/sh
docker volume create mariadb_data_tower
docker volume create mariadb_logs_tower
# docker volume 권한(설치 후, DB 접근 권한 이슈가 발생함)때문에 아래처럼, 수동으로 구성한다.
# 접근권한: chown -R 27:27 /tmp
docker run -v mariadb_data_tower:/tmp busybox chown -R 27:27 /tmp
docker run -v mariadb_logs_tower:/tmp busybox chown -R 27:27 /tmp
