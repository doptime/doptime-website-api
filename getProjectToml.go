package main

import (
	"bytes"

	"github.com/BurntSushi/toml"
	"github.com/doptime/doptime/api"
	"github.com/doptime/doptime/config"
	"github.com/doptime/doptime/rdsdb"
	"github.com/redis/go-redis/v9"
)

type ReqGetProjectToml struct {
	ProjectID string
	JwtID     string `validate:"required"`
}

var keyProjectToml = rdsdb.HashKey[string, string](rdsdb.Option.WithKey("ProjectToml"))

var APIGetProjectToml = api.Api(func(req *ReqGetProjectToml) (projectToml string, err error) {
	if projectToml, err = keyProjectToml.ConcatKey(req.JwtID).HGet(req.ProjectID); err != nil && err != redis.Nil {
		return "", err
	} else if err == nil {
		return projectToml, nil
	}
	//set w to memory writer
	desCfg := config.Configuration{
		ConfigUrl: "",
		Redis:     []*config.ConfigRedis{{Name: "default", Username: "your_redis_username", Password: "123456", Host: "localhost", Port: 6379, DB: 0}},
		Http:      config.ConfigHttp{CORES: "*", Port: 80, Path: "/", MaxBufferSize: 10485760},
		HttpRPC:   []*config.ApiSourceHttp{{Name: "doptime", UrlBase: "https://api.doptime.com", Jwt: ""}},
		Settings:  config.ConfigSettings{LogLevel: 1},
	}
	bytesBuffer := bytes.NewBuffer([]byte{})
	if err = toml.NewEncoder(bytesBuffer).Encode(desCfg); err != nil {
		return "", err
	}
	projectToml = bytesBuffer.String()

	return projectToml, nil
}).Func

type ReqSetProjectToml struct {
	ProjectID   string `validate:"required"`
	ProjectToml string `validate:"required"`
	JwtID       string `validate:"required"`
}

var APISetProjectToml = api.Api(func(req *ReqSetProjectToml) (result string, err error) {
	desCfg := &config.Configuration{
		ConfigUrl: "",
		Redis:     []*config.ConfigRedis{},
		Http:      config.ConfigHttp{CORES: "*", Port: 80, Path: "/", MaxBufferSize: 10485760},
		HttpRPC:   []*config.ApiSourceHttp{{Name: "doptime", UrlBase: "https://api.doptime.com", Jwt: ""}},
		Settings:  config.ConfigSettings{LogLevel: 1},
	}
	if _, err = toml.Decode(req.ProjectToml, desCfg); err != nil {
		return "", err
	}
	err = keyProjectToml.ConcatKey(req.JwtID).HSet(req.ProjectID, req.ProjectToml)
	return "", err
}).Func
