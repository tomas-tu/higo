# 指定基于的基础镜像
FROM scratch

# 维护者信息
MAINTAINER tomas7tu "tchenko.ok@gmail.com"

WORKDIR $GOPATH/src/github.com/tomas7tu/higo
COPY . $GOPATH/src/github.com/tomas7tu/higo

EXPOSE 8088
CMD ["./higo"]
