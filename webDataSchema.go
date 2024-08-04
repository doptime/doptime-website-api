package main

import "github.com/doptime/doptime/rdsdb"

type SubProjectIterator struct {
	ID        string
	CreateAt  int64
	EditAt    int64
	RootPath  string
	SkipDirs  string
	SkipFiles string
}

var KeySubProjectIterator = rdsdb.HashKey[string, *SubProjectIterator](rdsdb.WithRegisterWebData(true))

type ProjectInfo struct {
	ID       string
	Name     string
	EditAt   int64
	CreateAt int64
}

var KeyProjectInfo = rdsdb.HashKey[string, *ProjectInfo](rdsdb.WithRegisterWebData(true))
