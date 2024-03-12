package main

import (
	"fmt"
	"time"

	"github.com/yangkequn/goflow/api"
	_ "github.com/yangkequn/goflow/httpserve"
)

type ReqHello struct {
	HeaderIP        string
	HeaderUserAgent string
	Text            string
}

var ApiHello = api.New(func(req *ReqHello) (ret string, err error) {
	var response = fmt.Sprintf("Hello, IP:%s, UserAgent:%s, Text:%s", req.HeaderIP, req.HeaderUserAgent, req.Text)
	return response, nil

}, *api.Option.WithDataSource(""))

func main() {
	//forever sleep
	time.Sleep(time.Duration(1<<63 - 1))
}
