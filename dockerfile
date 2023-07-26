# 使用官方 Golang 镜像作为基础镜像，选择 latest 标签
FROM golang:latest

# 设置工作目录
WORKDIR /app

# 复制当前目录下的所有文件到容器的工作目录
COPY . .

# 编译 Go 应用
RUN go build -o main

# 暴露应用的端口号
EXPOSE 8080

# 设置容器启动时执行的命令
CMD ["./main"]
