# Copyright (c) 1999-2022 NAVER corp.
# All Rights Reserved.
#
# This software is the confidential and proprietary information of NAVER corp.
#
# You shall not disclose such Confidential Information and shall use it
# only in accordance with the terms of the license agreement that
# you entered into with NAVER
#
# @file   Makefile
#
.PHONY: all check build run push exec clean rmi

ifndef VERSION
$(error VERSION is not set)
endif

BUILD_OPTIONS = --no-cache --progress plain
NAME = ctower-httpd

all: build

build:
	DOCKER_BUILDKIT=1 \
	docker build $(BUILD_OPTIONS) \
		--build-arg OSS_BRANCH="main" \
		-t reg.navercorp.com/tower/httpd:$(VERSION) .

push:
	docker push reg.navercorp.com/tower/httpd:$(VERSION)

run:
	docker run --name $(NAME) -itd -p 8080:8080 reg.navercorp.com/tower/httpd:$(VERSION)

exec:
	docker exec -it $(NAME) bash

clean:
	docker rm -f $(NAME) || true

rmi:
	docker rmi reg.navercorp.com/tower/httpd:$(VERSION)

#vim: enc=utf-8 ts=4
