# 使用 golang 作为构建阶段
FROM golang AS builder
# 环境变量
ENV GO111MODULE=on \
        GOPROXY=https://goproxy.cn,direct \
        GOARCH=amd64 \
        GOOS=linux \
        CGO_ENABLED=0

# 移动到工作目录
WORKDIR /build
# 把代码复制到容器中
COPY . .
# 编译成可执行文件
RUN go build -o bluebell_app .

# 接下来创建一个小镜像
FROM debian:stretch-slim
COPY ./wait-for-it.sh /wait-for-it.sh
# 复制config.yaml文件 如果还有静态文件也复制静态文件
COPY ./conf /conf
# 把二进制文件从 builder 复制到 /
COPY --from=builder /build/bluebell_app /bluebell_app
# 声明服务端口
EXPOSE 8080
# 启动容器时运行的命令
# 在外面运行
# CMD [ "/bluebell_app" ]
