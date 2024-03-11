FROM golang:latest as build

# set env
ENV GO111MODULE=on

# turn on GO PROXY if you are in China
#ENV GOPROXY=https://goproxy.cn,direct

# set workdir
WORKDIR /go/release

RUN git clone -b master --single-branch https://github.com/yangkequn/goflow .
RUN go mod download

# copy local source file to container
COPY . .

# compile binary file
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o /go/release/goflowapp ./main.go
RUN chmod +x /go/release/goflowapp

# 准备包含 ca-certificates 的中间镜像
FROM alpine:latest as certs
RUN apk --update add ca-certificates

# create a new minimal image
FROM scratch as prod

# set timezone
#COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# copy binary file from build image
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/release/goflowapp /
CMD ["./goflowapp"]