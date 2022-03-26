FROM golang:1.16 AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build

COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct && go mod tidy
RUN go build -o httpserver .
FROM scratch
COPY --from=builder /build/httpserver /
EXPOSE 80
ENTRYPOINT ["/httpserver"]