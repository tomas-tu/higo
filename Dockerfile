# 指定基于的基础镜像
FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct

# 维护者信息
MAINTAINER tomas7tu "tchenko.ok@gmail.com"

WORKDIR $GOPATH/src/github.com/tomas7tu/higo
COPY . $GOPATH/src/github.com/tomas7tu/higo
RUN go build .

ENTRYPOINT ["./higo"]
