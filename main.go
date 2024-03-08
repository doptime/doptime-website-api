package main

import (
	"fmt"

	"github.com/yangkequn/goflow"
	"github.com/yangkequn/goflow/api"
	"github.com/yangkequn/goflow/config"
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
	goflow.Start()
	config.LoadConfig_FromEnv()
}
