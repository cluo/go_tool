# Go_tool
>This is a tool library for Golang.Dont't worry about not understant it!
>All comment writes by English,Ahaha~~ 

>Oh,I think some will be Chinese.

# Usage
```
go get -u -v github.com/hunterhug/go_tool
```

# Include
## Image 图像处理库 **(image deal library)**

```
package image

import (
	"testing"
	"fmt"
)

func TestImage(t *testing.T) {
	err := ThumbnailF2F("../data/image.png", "../data/image100*100.png", 100, 100)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = ScaleF2F("../data/image.png", "../data/image200.png", 200)
	if err != nil {
		fmt.Println(err.Error())
	}

	filename, err := RealImageName("../data/image.png")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("real filename"+filename)
	}
}

```

## Spider 爬虫封装库 **(spider library)**

```
package spider

import (
	"testing"
	"fmt"
)

func TestSpider(t *testing.T) {
	// a new spider
	spiders := Spider{}
	// method can be get and post
	spiders.Method = "get"
	// wait times,can zero
	spiders.Wait = 2
	// which url fetch
	spiders.Url = "http://www.goole.com"

	// log record
	//spiders.SetLogLevel("DEBUg")
	spiders.SetLogLevel("error")

	// a new header,default ua, no refer
	spiders.NewHeader(nil, "www.google.com", nil)

	// a proxy client
	proxy := "http://smart:smart2016@104.128.121.46:808"
	client, err := NewProxyClient(proxy)
	if err != nil {
		spiders.Log().Error(err.Error())
	}

	// set a client in a spider,if client no set,will use defalut client
	spiders.Client = client

	// go!fetch url --||
	body, err := spiders.Go()
	if err != nil {
		spiders.Log().Error(err)
	} else {
		// bytes get!
		fmt.Printf("%s", string(body))
	}

	// if filesize small than 500KB
	err = TooSortSizes(body, 500)

	spiders.Log().Error(err.Error())
}

```

## Dbs   数据库封装库 **(database library)**

Mysql

```
package mysql

import (
	"testing"
	"fmt"
)

/*

CREATE TABLE `51job_keyword` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `keyword` varchar(255) NOT NULL DEFAULT '',
  `address` varchar(255) NOT NULL DEFAULT '',
  `kind` varchar(255) NOT NULL DEFAULT '',
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL,
  `time51` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='关键字表';

*/

func TestMysql(t *testing.T) {

	config := MysqlConfig{
		Username:"root",
		Password:"6833066",
		Ip:"127.0.0.1",
		Port:"3306",
		Dbname:"51job",
	}

	db := New(config)

	db.Open()

	//'1', '教师', '潮州', '0', '2016-05-27 00:00:00', '2016-05-27 00:00:00', '204'
	sql:="INSERT INTO `51job_keyword`(`keyword`,`address`,`kind`) values(?,?,?)"

	num,err:=db.Insert(sql,"PHP","潮州",0)
	if err!=nil{
		fmt.Println(err.Error())
	}else{
		fmt.Printf("插入条数%d\n",num)
	}

	sql = "SELECT * FROM 51job_keyword where address=? and kind=? limit ?;"
	result, err := db.Select(sql, "潮州", 0, 6)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for row, v := range result {
			fmt.Printf("%v:%#v\n", row, v)
		}
	}
}

```

Cassandra

```
package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"testing"
)

/*　测试之前先填充一下cassandra语句
create keyspace example with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
create table example.tweet(timeline text, id UUID, text text, PRIMARY KEY(id));
create index on example.tweet(timeline);
*/
func TestCdb(t *testing.T) {
	//先构造一个字符串数组连接
	host := []string{"192.168.11.74"}

	//指定cassandra keyspace，类似于mysql中的db
	keyword := "example"

	//连接
	cdb := NewCdb(host, keyword)


	//构造插入语句
	insertsql := cdb.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
		"me", gocql.TimeUUID(), "hello world")

	//执行,查看Exec方法,可知执行后已经关闭
	if err := insertsql.Exec(); err != nil {
		log.Fatal(err)
	}

	//构造查找语句
	querysql := cdb.Query(`SELECT "id", text FROM "tweet"`)

	//执行
	iter := querysql.Iter()

	//定义字节数组
	var id gocql.UUID
	var text string

	//循环取值，需要手工，无法再封装
	for iter.Scan(&id, &text) {
		fmt.Printf("Tweet:%v,%v\n", id, text)
	}
	//这个需要关闭
	if err := iter.Close(); err != nil {
		fmt.Printf("%v", err)
	}


}

```

Hbase

```
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
```
## Http 网络库　**(network library)**

## Util 文件/时间等杂项库 **(some small library)**

```
package util

import (
	"testing"
	"fmt"
)

func TestUtil(t *testing.T) {
	i := 2
	if ("2" == IS(i)) {
		fmt.Println("int to string")
	}

	v, err := SI("e2")
	if (err == nil&&v == i) {
		fmt.Println("string to int")
	} else {
		fmt.Println(err.Error())
	}

	fmt.Println(CurDir())

	err = MakeDir("../data")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("dir already exist")
	}

	err = MakeDirByFile("../data/testutil.txt")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("dir already exist")
	}

	err = SaveTofile("../data/testutil.txt", []byte("testutil"))
	if err != nil {
		fmt.Println(err.Error())
	}
}

```

# How to use
>You all can read the test golang file.And I recomment use IDE **pycharm** which python language use,
can also install The Go plugin.

# Author
>一只尼玛

>My website:http://www.lenggirl.com
 
><p>Updating...