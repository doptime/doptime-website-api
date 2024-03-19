package main

import (
	"fmt"
	"time"

	"github.com/doptime/doptime/api"
	_ "github.com/doptime/doptime/httpserve"
)

type ReqHello struct {
	HeaderRemoteAddr string
	HeaderUserAgent  string
	Text             string
}

var ApiHello = api.New(func(req *ReqHello) (ret string, err error) {
	var response = fmt.Sprintf("Hello, IP:%s, UserAgent:%s, Text:%s", req.HeaderRemoteAddr, req.HeaderUserAgent, req.Text)
	return response, nil

}, *api.Option.WithDataSource(""))

func main() {
	//forever sleep
	time.Sleep(time.Duration(1<<63 - 1))
}
