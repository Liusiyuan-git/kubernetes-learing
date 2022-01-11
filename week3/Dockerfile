FROM golang:1.15 AS build
RUN git clone https://github.com/Liusiyuan-git/kubernetes-learing.git
WORKDIR kubernetes-learing/week3
RUN go mod init week3 && go build -o httpserver week3
ENTRYPOINT ["/go/kubernetes-learing/week3/httpserver"]