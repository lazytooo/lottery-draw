FROM golang:alpine3.14 AS build-env


RUN mkdir -p /lottery_draw
WORKDIR /lottery_draw/

COPY main /lottery_draw/


RUN chmod 777 main

ENV DB_ADDR=""
ENV DB_PORT=""
ENV DB_NAME=""
ENV DB_USER=""
ENV DB_PASS=""
ENV DB_MAX_IDLE=""
ENV DB_MAX_OPEN=""
ENV REDIS_ADDR=""
ENV REDIS_PASS=""
ENV WS_ADDR=""

CMD ["/lottery_draw/main"]

