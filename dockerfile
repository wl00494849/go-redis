FROM golang
WORKDIR /go/src/go-redis
ADD . /go/src/go-redis

RUN apt-get update \
    && apt install -y git

RUN cd /go/src/go-redis \
    && go get github.com/go-sql-driver/mysql \
    && go get github.com/rs/cors \
    && go get github.com/gin-gonic/gin \
    && go get github.com/go-redis/redis/v8 \
    && go build

RUN mv go-redis /go/src/redis \
    && cd /go/src \
    && rm -r go-redis

EXPOSE 7788
ENTRYPOINT /go/src/redis