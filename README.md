# Go_tool
>This is a tool library for Golang.Dont't worry about not understant it!
>All comment writes by English,Ahaha~~ 

>Oh,I think some will be Chinese.

# Usage
```
go get -u -v github.com/hunterhug/go_tool
go get -v github.com/hunterhug/go_image
go get -v github.com/hunterhug/go-hbase
go get -v gopkg.in/redis.v4
go get -v github.com/gocql/gocql
```

# Include
## Image 图像处理库 **(image deal library)**

```
package image

import (
	"testing"
)

func TestImage(t *testing.T) {

	// Scale a image file by cuting 100*100
	err := ThumbnailF2F("../data/image.png", "../data/image100-100.png", 100, 100)
	if err != nil {
		t.Error("Test ThumbnailF2F:" + err.Error())
	}

	// Scale a image file by cuting width:200 (Equal scaling)
	err = ScaleF2F("../data/image.png", "../data/image200.png", 200)
	if err != nil {
		t.Error("Test ScaleF2F:" + err.Error())
	}

	// File Real name
	filename, err := RealImageName("../data/image.png")
	if err != nil {
		t.Error("Test RealImageName:" + err.Error())
	} else {
		t.Log("Test RealImageName::real filename" + filename)
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
)

/*

CREATE TABLE IF NOT EXISTS `51job_keyword` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `keyword` varchar(255) NOT NULL DEFAULT '',
  `address` varchar(255) NOT NULL DEFAULT '',
  `kind` varchar(255) NOT NULL DEFAULT '',
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  `time51` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='关键字表';

*/

func TestMysql(t *testing.T) {

	// mysql config
	config := MysqlConfig{
		Username: "root",
		Password: "smart2016",
		Ip:       "127.0.0.1",
		Port:     "3306",
		Dbname:   "51job",
	}

	// a new db connection
	db := New(config)

	// open connection
	db.Open()

	// create sql
	sql := `
  CREATE TABLE IF NOT EXISTS 51job.51job_keyword (
  id int(11) NOT NULL AUTO_INCREMENT,
  keyword varchar(255) NOT NULL DEFAULT '',
  address varchar(255) NOT NULL DEFAULT '',
  kind varchar(255) NOT NULL DEFAULT '',
  created datetime DEFAULT NULL,
  updated datetime DEFAULT NULL,
  time51 int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='关键字表';`

	// create
	inum, err := db.Create(sql)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Logf("create number:%d\n", inum)
	}

	// insert sql
	//'1', '教师', '潮州', '0', '2016-05-27 00:00:00', '2016-05-27 00:00:00', '204'
	sql = "INSERT INTO `51job_keyword`(`keyword`,`address`,`kind`) values(?,?,?)"

	// insert
	num, err := db.Insert(sql, "PHP", "潮州", 0)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Logf("insert number:%d\n", num)
	}

	// select sql
	sql = "SELECT * FROM 51job_keyword where address=? and kind=? limit ?;"

	// select
	result, err := db.Select(sql, "潮州", 0, 6)
	if err != nil {
		t.Error(err.Error())
	} else {
		// values
		for row, v := range result {
			t.Logf("%v:%#v\n", row, v)
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
before test,To do bellow

create_namespace 'hunterhug'
create 'hunterhug:info','colfamily',{SPLITALGO => 'HexStringSplit',NUMREGIONS => 40}
*/


func TestHbase(t1 *testing.T) {
	var (
		//格式化时间
		// Time format
		fs = "2006-01-02 15:04:05"

		//hbase配置
		//hbase config
		config = HbaseConfig{
			Zkport: "2181",//zk
			Zkquorum: "192.168.11.73",
		}
		//创建一个客户端
		// new client
		clients = New(config)
	)

	//打开客户端
	//open connection
	clients.Open()

	//放主键
	//set a rowkey
	rowkey := "rowkey"

	//命名空间
	//set namespace
	hbasenamespace := "hunterhug"

	//表名
	//set tablename
	hbasetable := "info"

	//列族
	//set colfamily
	hbasefamily := "colfamily"

	//set col
	hbasecol := "go"
	hbasecol1 := "die"

	//放值，第一列
	// put a value to  hunterhug:info colfamily:go
	put := hbase.CreateNewPut([]byte(rowkey))
	put.AddStringValue(hbasefamily, hbasecol, "value")
	_, err := clients.Client.Put(hbasenamespace + ":" + hbasetable, put)

	if err != nil {
		t1.Logf(err.Error())
	}


	//第二列
	// the same,put another value
	put.AddStringValue(hbasefamily, hbasecol1, "value1")
	_, err = clients.Client.Put(hbasenamespace + ":" + hbasetable, put)

	if err != nil {
		t1.Logf(err.Error())
	}

	//获取一行
	//get value by rowkey
	result, err := clients.GetResult(hbasenamespace + ":" + hbasetable, rowkey)
	if err != nil {
		t1.Logf(err.Error())
	}
	if v, exist := result.Columns[hbasefamily + ":" + hbasecol]; exist {
		fmt.Printf("found row:%v,v:%v\n",rowkey,v.Value.String())
	}

	//扫描
	// create a scan object,so scan can be close because client is inside in it
	scan := clients.Client.Scan(hbasenamespace + ":" + hbasetable)

	//过滤两列，只拿
	// just scan one col
	err = scan.AddString(hbasefamily + ":" + hbasecol)

	if err != nil {
		panic(err)
	}

	// no!we want two col
	err = scan.AddString(hbasefamily + ":" + hbasecol1)
	if err != nil {
		panic(err)
	}

	//过滤时间戳，只拿时间戳在这段范围的合数据
	// and we are just want before 720h's data
	var t = time.Now().Add(-1 * 720 * time.Hour * 1)

	fmt.Printf("Get Data which record before %v\n", t.Format(fs))

	// Time filter
	scan.SetTimeRangeFrom(t)

	var v1, v2 string

	//循环处理值
	// get scan value
	scan.Map(func(r *hbase.ResultRow) {
		// if value exist,v1!
		if v, exist := r.Columns[hbasefamily + ":" + hbasecol]; exist {
			v1 = v.Value.String()
		}
		if v, exist := r.Columns[hbasefamily + ":" + hbasecol1]; exist {
			v2 = v.Value.String()
		}
		// output
		fmt.Printf("row:%s\t  %s:%s \t %s:%s \t", r.Row.String(), hbasecol, v1, hbasecol1, v2)
	})
}
```

## Http 网络库　**(network library)**

## Util 文件/时间等杂项库 **(some small library such as times and file)**

```
package util

import (
	"testing"
)

func TestUtil(t *testing.T) {
	// test int to string
	i := 2
	if "2" == IS(i) {
		t.Log("Test IS:int to string")
	}

	// test string to int
	v, err := SI("e2")
	if err == nil && v == i {
		t.Log("Test SI:string to int")
	} else {
		t.Log("Test SI:" + err.Error())
	}

	// caller dir
	t.Log("Test CurDir:"+CurDir())

	// create dir
	err = MakeDir("../data")
	if err != nil {
		t.Log("Test MakeDir:" + err.Error())
	} else {
		t.Log("Test MakeDir:dir already exist")
	}

	// create dir by filename
	filename := "../data/testutil.txt"
	err = MakeDirByFile(filename)
	if err != nil {
		t.Log("Test MakeDirByFile:" + err.Error())
	} else {
		t.Log("Test MakeDirByFile: dir already exist")
	}

	// save bytes into file
	err = SaveToFile(filename, []byte("testutil"))
	if err != nil {
		t.Log("Test SaveToFile" + err.Error())
	}

	// read bytes from file
	filebytes, err := ReadfromFile(filename)
	if err != nil {
		t.Error("Test ReadfromFile:" + err.Error())
	} else {
		t.Log("Test ReadfromFile:" + string(filebytes))
	}
}

```

# How to use
>You all can read the test golang file.And I recomment use IDE **pycharm** which python language use,
can also install The Go plugin.

# Author
>一只尼玛

My website:[http://www.lenggirl.com](http://www.lenggirl.com)

My Twitter:[https://twitter.com/hunterhug_](https://twitter.com/hunterhug_)

My Weibo:[http://weibo.com/hunterhug](http://weibo.com/hunterhug)

><p>Updating...