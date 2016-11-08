package myhbase

import (
	"strings"
	"github.com/hunterhug/go-hbase"
	"errors"
)

//hbase配置结构体,modify by java
// a hbase config
/*
  "master": "192.168.11.73:60000",
  "zkport": "2181",
  "zkquorum": "192.168.11.73,192.168.11.74"
*/
type HbaseConfig struct {
	//zookeeper port
	Zkport   string
	//zk ip
	Zkquorum string
}

// a hbase client
type Hbase struct {
	Config HbaseConfig
	Client *hbase.Client
}

// create a new hbase client
func New(config HbaseConfig) Hbase {
	return Hbase{Config:config}
}

//太坑了，坑啊，表名前缀不能太相似！！such as aaaaaaaaaaaaaaaaa,aaaaaaaaaaaaaaaaaab这样可能会把数据发送到另外的表
// init hbase connection,the table name can't be have a long same prefix
func (db *Hbase)Open() {

	// get hbase config
	config := db.Config

	// many zkquorum but port is the same, will be fix
	//Todo

	zkhosts := strings.Split(config.Zkquorum, ",")

	// join
	for i, _ := range zkhosts {
		zkhosts[i] = zkhosts[i] + ":" + config.Zkport
	}
	// zkroot,where to find hbase position
	zkroot := "/hbase"

	//start
	client := hbase.NewClient(zkhosts, zkroot)
	//client.SetLogLevel("DEBUG")
	db.Client = client
}

//获取结果
// Get a result by hbase rowkey
func (db *Hbase)GetResult(table string, rowkey string) (result *hbase.ResultRow, err error) {
	get := hbase.CreateNewGet([]byte(rowkey))
	result, err = db.Client.Get(table, get)
	if (rowkey != result.Row.String()) {
		err = errors.New("No rowkey")
	}
	return
}
