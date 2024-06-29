package main

import "github.com/doptime/doptime/rdsdb"

type SubProjectIterator struct {
	ID           string
	CreateAt     int64
	EditAt       int64
	RootPath     string
	PathsIgnored string
	FilesIgnored string
}

var KeySubProjectIterator = rdsdb.HashKey[string, *SubProjectIterator](rdsdb.Option.WithRegisterWebDataSchema())

type ProjectInfo struct {
	ID       string
	Name     string
	EditTime int64
}

var KeyProjectInfo = rdsdb.HashKey[string, *ProjectInfo](rdsdb.Option.WithRegisterWebDataSchema())
