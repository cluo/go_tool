# 一.Go_tool
>This is a tool library for Golang.Dont't worry about not understant it!
>All comment writes by English,Ahaha~~ 

>Oh,I think some will be Chinese.

# 二.Usage
```
go get -u -v github.com/hunterhug/go_tool
go get -v github.com/hunterhug/go_image
go get -v github.com/hunterhug/go-hbase
go get -v gopkg.in/redis.v4
go get -v github.com/gocql/gocql
go get -v golang.org/x/net/context
```

# 三.Include

## 1.Image 图像处理库 **(image deal library)**

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

## 2.Spider 爬虫封装库 **(spider library)**

```
package spider

import (
	"testing"
	"fmt"
)

func TestSpider(t *testing.T) {

	// global log record
	//SetLogLevel("DEBUg")
	SetLogLevel("debug")
	// GLOBAL TIMEOUT
	SetGlobalTimeout(3)

	// a new spider without proxy
	// NewSpider(nil)
	proxy := "http://smart:smart2016@104.128.121.46:808"
	spiders,err := NewSpider(proxy)
	if err!=nil{
		panic(err)
	}
	// method can be get and post
	spiders.Method = "get"
	// wait times,can zero
	spiders.Wait = 2
	// which url fetch
	spiders.Url = "http://www.goole.com"

	// a new header,default ua, no refer
	spiders.NewHeader(nil, "www.google.com", nil)


	// go!fetch url --||
	body, err := spiders.Go()
	if err != nil {
		Log().Error(err)
	} else {
		// bytes get!
		fmt.Printf("%s", string(body))
	}

	Log().Debugf("%#v",spiders)

	// if filesize small than 500KB
	err = TooSortSizes(body, 500)

	Log().Error(err.Error())
}

```

## 3.Dbs   数据库封装库 **(database library)**

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
	db.Open(2000,1000)

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

Redis

```
package myredis

import (
	"testing"
)

func TestRedis(t *testing.T) {
	config := RedisConfig{}
	config.DB = 0
	config.Host = "127.0.0.1:6379"
	config.Password = "smart2016" // no password set ""

	client, err := NewRedis(config) // new redis client
	if err != nil {
		panic(err)
	}

	// set key==value
	err = client.Set("key", "value", 0)
	if err != nil {
		t.Error(err.Error())
	}

	// get key
	val, err := client.Get("key")
	if err != nil {
		panic(err)
	} else {
		t.Log("Redis value:" + val)
	}

	// push test,pust pool with b value
	num, err := client.Lpush("pool", "b")
	if err != nil {
		t.Error(err.Error())
	}

	// total length of list
	t.Log(num)

	// pushx test,will be error if not exist pool10
	num, err = client.Lpushx("pool10", "b")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(num)

	// len test
	t.Log(client.Llen("pool"))

	// pop test
	pops, e := client.Lpop("pool2")
	if e!=nil {
		t.Logf("%v,%v", pops, e)
	}

	// bpop test
	bpops, e := client.Blpop(0,"pool1","pool1")
	t.Logf("%#v,%v", bpops, e)

	//rpoplpush test POOL1 empty so will be redis.nil
	rpoplpush,e:=client.Rpoplpush("POOL1","pool1")
	t.Logf("%#v,%v", rpoplpush, e)

	//brpoplpush test POOL1 empty so will be redis.nil if timeout but zero set is wait a long time
	brpoplpush,e:=client.Brpoplpush("POOL1","pool1",15)
	t.Logf("%#v,%v", brpoplpush, e)
}

```

Cassandra

```
package cassandra

import (
	"github.com/gocql/gocql"
	"testing"
)

/*　
测试之前先填充一下cassandra语句
Before test:

create keyspace example with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
create table example.tweet(timeline text, id UUID, text text, PRIMARY KEY(id));
create index on example.tweet(timeline);
*/
func TestCdb(t *testing.T) {
	//先构造一个字符串数组连接
	// cassandra host
	host := []string{"192.168.11.74"}

	//指定cassandra keyspace，类似于mysql中的db
	// cassandra keyspace like mysql db
	keyword := "example"

	//连接
	// connect a cdb
	cdb := NewCdb(host, keyword)


	//构造插入语句
	// insert sql!
	insertsql := cdb.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
		"me", gocql.TimeUUID(), "hello world")

	//执行,查看Exec方法,可知执行后已经关闭
	// exec a insert operation
	if err := insertsql.Exec(); err != nil {
		t.Fatal(err)
	}

	//构造查找语句
	// query sql
	querysql := cdb.Query(`SELECT "id", text FROM "tweet"`)

	//执行
	// done it !
	iter := querysql.Iter()

	//定义字节数组
	// id uuid in cassandra
	var id gocql.UUID
	var text string

	//循环取值，需要手工，无法再封装
	// take value just like that i can't wrap it again, and no need
	for iter.Scan(&id, &text) {
		t.Logf("Tweet:%v,%v\n", id, text)
	}

	//这个需要关闭
	// should be close
	if err := iter.Close(); err != nil {
		t.Logf("%v", err)
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


## 4.Util 文件/时间等杂项库 **(some small library such as times and file)**

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

	// times format
	t.Log(TodayString(3))

	// file exist?
	t.Logf("%v",FileExist("../r.txt"))

	// find the go file in some dir
	filenames,err:=ListDir(`G:\smartdogo\src\github.com\hunterhug\go_tool`,"go")
	t.Logf("%v:%v",filenames,err)

	// devide a string list into severy string list
	stringlist:=[]string{"2","3","4","5","4"}
	num:=3
	result,err:=DevideStringList(stringlist,num)
	if err!=nil{
		t.Error(err)
	}else{
		t.Logf("%#v",result)
	}

	// now secord times from January 1, 1970 UTC.
	secord:=GetSecordTimes()
	t.Log(secord)

	// now date string by secord
	timestring:=GetSecord2DateTimes(secord)
	t.Log(timestring)

	// change back
	t.Log(GetDateTimes2Secord(timestring))
}

```

# 四.How to use
>You all can read the test golang file.And I recomment use IDE **pycharm** which python language use,
can also install The Go plugin.

# 五.Author
>太阳萌飞了

My website:[http://www.lenggirl.com](http://www.lenggirl.com)

My Twitter:[https://twitter.com/hunterhug_](https://twitter.com/hunterhug_)

My Weibo:[http://weibo.com/hunterhug](http://weibo.com/hunterhug)

><p>Updating...