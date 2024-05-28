FROM golang:1.22-alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 下载依赖信息
RUN go mod download

# 将我们的代码编译成二进制可执行文件 bluebell_app
RUN go build -o bluebell_app .

###################
# 接下来创建一个小镜像
###################
FROM scratch

# 从builder镜像中把静态文件拷贝到当前目录
COPY ./templates /templates
COPY ./static /static

# 从builder镜像中把配置文件拷贝到当前目录
COPY ./conf /conf

#设置一个环境变量，让docker中的程序在docker容器中寻找
ENV CONFIG_PATH=/conf/config.yaml

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/bluebell_app /

# 需要运行的命令
ENTRYPOINT ["/bluebell_app", "conf/config.yaml"]
