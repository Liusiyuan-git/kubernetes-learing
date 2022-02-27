FROM golang:1.17 AS builder
ENV GO111MODULE=off \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct && go mod init httpserver && go mod tidy
RUN go build -o httpserver .
FROM scratch
COPY --from=builder /build/httpserver /
EXPOSE 8080
ENTRYPOINT ["/httpserver"]