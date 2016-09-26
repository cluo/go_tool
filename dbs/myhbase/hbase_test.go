package myhbase

import (
	"github.com/hunterhug/go-hbase"
	"testing"
	"fmt"
	"time"
)

/*
测试之前
create_namespace 'hunterhug'
create 'hunterhug:info','colfamily',{SPLITALGO => 'HexStringSplit',NUMREGIONS => 40}
*/


func TestHbase(t1 *testing.T) {
	var (
		//格式化时间
		fs = "2006-01-02 15:04:05"

		//hbase配置
		config = HbaseConfig{
			Zkport: "2181",
			Zkquorum: "192.168.11.73",
		}
		//创建一个客户端
		clients = New(config)
	)

	//打开客户端
	clients.Open()

	//放主键
	rowkey := "rowkey"

	//命名空间
	hbasenamespace := "hunterhug"

	//表名
	hbasetable := "info"

	//列族
	hbasefamily := "colfamily"

	//列
	hbasecol := "go"
	hbasecol1 := "die"

	//放值，第一列
	put := hbase.CreateNewPut([]byte(rowkey))
	put.AddStringValue(hbasefamily, hbasecol, "value")
	_, err := clients.Client.Put(hbasenamespace + ":" + hbasetable, put)

	if err != nil {
		fmt.Println(err.Error())
	}
	//第二列
	put.AddStringValue(hbasefamily, hbasecol1, "value1")
	_, err = clients.Client.Put(hbasenamespace + ":" + hbasetable, put)

	if err != nil {
		fmt.Println(err.Error())
	}

	//获取一行
	result, err := clients.GetResult(hbasenamespace + ":" + hbasetable, rowkey)
	if err != nil {
		fmt.Println(err.Error())
	}
	if v, exist := result.Columns[hbasefamily + ":" + hbasecol]; exist {
		fmt.Printf("找到行row:%v,v:%v\n",rowkey,v.Value.String())
	}

	//扫描
	scan := clients.Client.Scan(hbasenamespace + ":" + hbasetable)

	//过滤两列，只拿
	err = scan.AddString(hbasefamily + ":" + hbasecol)

	if err != nil {
		panic(err)
	}
	err = scan.AddString(hbasefamily + ":" + hbasecol1)
	if err != nil {
		panic(err)
	}

	//过滤时间戳，只拿时间戳在这段范围的合数据
	var t = time.Now().Add(-1 * 720 * time.Hour * 1)
	fmt.Printf("寻找时间在%v之后的记录\n", t.Format(fs))

	scan.SetTimeRangeFrom(t)

	var v1, v2 string

	//循环处理值
	scan.Map(func(r *hbase.ResultRow) {

		if v, exist := r.Columns[hbasefamily + ":" + hbasecol]; exist {
			v1 = v.Value.String()
		}
		if v, exist := r.Columns[hbasefamily + ":" + hbasecol1]; exist {
			v2 = v.Value.String()
		}
		fmt.Printf("row:%s\t  %s:%s \t %s:%s \t", r.Row.String(), hbasecol, v1, hbasecol1, v2)
	})
}