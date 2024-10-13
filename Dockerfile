FROM golang:1.23.1-alpine3.20 AS builder

COPY . /app

WORKDIR /app

RUN export GOPATH=${HOME}/go && \
    export GOBIN=$HOME/go/bin && \
    export PATH=$PATH:$GOROOT/bin:$GOBIN && \
    apk update && apk add --no-cache make protobuf-dev && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    sh pkg/proto/v1/compile_proto.sh && \
    go mod tidy  && \
    CGO_ENABLED=0 GOOS=linux go build -o /go_grpc_demo cmd/main.go


FROM scratch

COPY --from=builder /etc/ssl /etc/ssl

COPY --from=builder /go_grpc_demo /go_grpc_demo

EXPOSE 8082

CMD ["/go_grpc_demo"]



