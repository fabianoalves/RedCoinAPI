
FROM golang:latest

WORKDIR /go/src/github.com/rteles86/RedCoinApi/redcoin

ADD . .

WORKDIR /go/src/github.com/rteles86/RedCoinApi/redcoin

RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/go-redis/redis
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/mux

RUN go build -o userServer

ENTRYPOINT ./userServer

EXPOSE 2801

EXPOSE 1805

EXPOSE 6379