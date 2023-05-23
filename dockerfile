FROM golang:alpine as build

# 设置环境变量
ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64 \
  GOPROXY="https://goproxy.cn,direct"

WORKDIR /go/release

ADD . .

# 编译
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o app main.go router_.go touter.go

# 运行：使用scratch作为基础镜像
FROM scratch as prod

# 在build阶段复制时区到
COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
# 在build阶段复制可执行的go二进制文件app
COPY --from=build /go/release/app /

# 启动服务
CMD ["/app"]