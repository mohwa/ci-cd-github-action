# https://velog.io/@seheon99/Dockerfile-%EC%9E%91%EC%84%B1-%EB%B0%A9%EB%B2%95-12
# syntax=docker/dockerfile:1.2
# -------------------------------------- #
# Stage 1 : Build Client
# -------------------------------------- #
ARG USER_NAME=app
ARG GITHUB_BRANCH=main

ARG HOME=/home/$USER_NAME
ARG MY_APP=$HOME/my-app

FROM node:16.14.2 AS vue-builder

# FROM 안에서 재 선언한다.
# https://stackoverflow.com/questions/53681522/share-variable-in-multi-stage-dockerfile-arg-before-from-not-substituted
ARG USER_NAME
ARG HOME
ARG MY_APP

# app 유저를 추가한다.(/home/app 폴더도 같이 생성된다)

# /bin/false 에 대한 글
# https://faq.hostway.co.kr/Linux_ETC/1624
RUN useradd --user-group --create-home --shell /bin/false $USER_NAME

# my_app 디렉토리를 추가한다.
# app 유저 권한을 $HOME 디렉토리(/home/app)의 하위에 모두 추가한다.
RUN mkdir -p $MY_APP && \
    chown -R $USER_NAME:$USER_NAME $HOME/*

WORKDIR $MY_APP

# RUN, CMD, ENTRYPOINT 명령이 수행될 유저를 설정한다.
# http://pyrasis.com/book/DockerForTheReallyImpatient/Chapter07/12
USER $USER_NAME

# 호스트와 연결될 포트 번호를 설정한다.
EXPOSE 3000

RUN git clone -b $GITHUB_BRANCH https://git@github.com/mohwa/ci-cd-github-action $MY_APP && \
    cd $MY_APP/client && \
    npm i && \
    npm run build

# -------------------------------------- #
# Stage 2 : Serving                      #
# -------------------------------------- #
FROM nginx as nginx

LABEL maintainer="mohwa <yanione2@gmail.com>"
LABEL description="Dockerfile for CI/CD Test"

ARG USER_NAME
ARG HOME
ARG MY_APP
ARG DEFAULT_CONF="default.conf"

EXPOSE 5000

COPY $DEFAULT_CONF /etc/nginx/conf.d/default.conf
COPY --from=vue-builder --chown=$USER_NAME:$USER_NAME $MY_APP/dist /usr/share/nginx/html

# 반드시 실행되는 명령이다.
# ENTRYPOINT ["${HOME}/apache/bin/apachectl"]
# docker run 시, 변경할 수 있는 인자들을 정의한다.
# 즉 docker run 시, 인자를 정의하지않았다면, 아래 정의된 기본 인자들로 apachectl 명령이 수행된다
# CMD ["-k", "start", "-D", "FOREGROUND"]
