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

	//执行,查看Ｅxec方法,可知执行后已经关闭
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

