FROM golang:1.19

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

# 移动到工作目录：/build

WORKDIR $GOPATH/src/gin-web
COPY . $GOPATH/src/gin-web
RUN go build .


# 声明服务端口
EXPOSE 3001

# 将代码复制到容器中
COPY . .

ENTRYPOINT ["./gin-web"]


