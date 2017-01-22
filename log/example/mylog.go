package main

import (
	"github.com/hunterhug/go_tool/util"
	"sunteng/commons/log"
)

func main() {
	log.New(util.CurDir() + "/log.json")
	logger := log.AmazonListLog
	logger.Error("hello logger sunteng/commons/log/a")
	logger.Warn("hello logger sunteng/commons/log/a")
	logger.Notice("hello logger sunteng/commons/log/a")
	logger.Log("hello logger sunteng/commons/log/a")
	//看不到debug信息
	logger.Debug("hello logger sunteng/commons/log/a")

	logger = log.Get("dayasin")
	logger.Log("hello logger sunteng/commons/log/a/b")

	logger = log.Get("dayip")
	logger.Debug("hello logger sunteng/commons/log/b")
	//看不到log信息
	logger.Log("hello logger sunteng/commons/log/b")
	logger.Error("hello logger sunteng/commons/log/b")
}
