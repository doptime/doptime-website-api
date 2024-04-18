package main

import (
	"fmt"
	"math/big"
	"time"

	"github.com/doptime/doptime/api"
	"github.com/doptime/doptime/config"
	"github.com/doptime/doptime/httpserve"
)

type ReqHello struct {
	HeaderRemoteAddr string
	HeaderUserAgent  string
	Text             string                 `validate:"required,min=10,max=10000"`
	Other            map[string]interface{} `mapstructure:",remain"`
}

var ApiHello = api.Api(func(req *ReqHello) (ret string, err error) {
	var response = fmt.Sprintf("Hello, IP:%s, UserAgent:%s, Text:%v", req.HeaderRemoteAddr, req.HeaderUserAgent, req.Text)
	return response, nil

})

func main() {
	config.LoadConfig_FromWeb()
	httpserve.Debug()
	bigint := big.NewInt(0)
	bigint.Text(62)
	httpserve.Debug()
	//forever sleep
	time.Sleep(time.Duration(1<<63 - 1))
}

var ApiRpcOverHttp = api.RpcOverHttp[*ReqHello, string]().HookParamEnhancer(func(req *ReqHello) (r *ReqHello, e error) {
	return req, nil
}).HookResultSaver(func(req *ReqHello, ret string) (e error) {
	//save result here if you need
	return nil
}).HookResponseModifier(func(req *ReqHello, ret string) (retvalue interface{}, e error) {
	//furthur modify the response value to the client here
	return ret, nil
}).Func
