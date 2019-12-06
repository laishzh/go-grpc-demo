FROM golang:1.13-alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk add git

RUN go get -v -u github.com/laishzh/go-grpc-demo/go_grpc_demo_server \
    && go install github.com/laishzh/go-grpc-demo/go_grpc_demo_server

EXPOSE 50051

CMD [ "go_grpc_demo_server" ]