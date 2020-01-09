FROM golang:latest
WORKDIR $GOPATH/src/github.com/shorturl
#将服务器的go工程代码加入到docker容器中
ADD . $GOPATH/src/github.com/shorturl
#go构建可执行文件
RUN go build .
#暴露端口
EXPOSE 8000
#最终运行docker的命令
ENTRYPOINT  ["./shorturl"]