FROM golang:alpine3.13

RUN apk update \
    && apk add tzdata git \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

WORKDIR $GOPATH/src/github.com/coomp/ccs-app
COPY . $GOPATH/src/github.com/coomp/ccs-app
RUN go build .
EXPOSE 8090
ENTRYPOINT ["./ccs-app"]
