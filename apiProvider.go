package main

// import (
// 	"github.com/doptime/doptime/api"
// 	"github.com/doptime/doptime/rdsdb"
// )

// type ReqGetApiProvider struct {
// 	ApiName string `validate:"required"`
// }

// var keyApiProvider = rdsdb.HashKey[string, *api.ApiProvider](rdsdb.Option.WithKey("ApiProvider"))
// var apiGetApiProvider = api.Api(func(req *ReqGetApiProvider) (ret string, err error) {
// 	return keyApiProvider.HScan(req.ApiName)
// }).Func

// func (am *APIManager) AddAPI(api API) error {
// 	// 添加API逻辑
// }
