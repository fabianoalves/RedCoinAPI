
FROM golang:latest

WORKDIR /go/src/github.com/rteles86/RedCoinApi/API

ADD . .

WORKDIR /go/src/github.com/rteles86/RedCoinApi/API

RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/go-redis/redis
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/mux
RUN go get gopkg.in/go-playground/validator.v10
RUN go get golang.org/x/crypto/bcrypt

RUN go build -o userServer

ENTRYPOINT ./userServer

EXPOSE 2801

EXPOSE 1805

EXPOSE 6379