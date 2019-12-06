FROM golang:1.13-alpine AS compiler

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk add git

RUN go get -v -u github.com/laishzh/go-grpc-demo/go_grpc_demo_client \
    && go install github.com/laishzh/go-grpc-demo/go_grpc_demo_client
RUN cp "${GOPATH}/bin/go_grpc_demo_client" /tmp/

FROM alpine
COPY --from=compiler /tmp/go_grpc_demo_client /usr/bin/go_grpc_demo_client

CMD [ "go_grpc_demo_client" ]