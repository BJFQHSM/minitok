FROM golang:1.18
MAINTAINER codingdog1412
WORKDIR /opt/tiktok

COPY go.mod go.sum .
RUN go env -w GOPROXY="https://goproxy.cn, direct" && go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/biz ./cmd/biz/*.go
RUN go build -v -o /usr/local/bin/auth ./cmd/auth/*.go
RUN go build -v -o /usr/local/bin/api ./cmd/api/*.go

RUN chmod -R 777 .
CMD ["/opt/tiktok/build.sh"]