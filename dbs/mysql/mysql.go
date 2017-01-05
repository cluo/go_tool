/*
 * Created by 一只尼玛 on 2016/8/12.
 * 功能： Mysql dbs
 *
 */
package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	//"fmt"
	"fmt"
)

// Mysql config
type MysqlConfig struct {
	Username string
	Password string
	Ip       string
	Port     string
	Dbname   string
}

// a client
type Mysql struct {
	Config MysqlConfig
	Client *sql.DB
}

func New(config MysqlConfig) Mysql {
	return Mysql{Config: config}
}

//插入数据
//Insert Data
func (db *Mysql) Insert(prestring string, parm ...interface{}) (int64, error) {
	stmt, err := db.Client.Prepare(prestring)
	if err != nil {
		//log.Println(err)
		return 0, err
	}
	R, err := stmt.Exec(parm...)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	num, err := R.RowsAffected()
	return num, err

}

// Create table
func (db *Mysql) Create(prestring string, parm ...interface{}) (int64, error) {
	stmt, err := db.Client.Prepare(prestring)
	if err != nil {
		//log.Println(err)
		return 0, err
	}
	R, err := stmt.Exec(parm...)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	num, err := R.RowsAffected()
	return num, err

}

// create database
func (dbconfig MysqlConfig) CreateDb() (int64, error) {
	dbname := dbconfig.Dbname
	sql := fmt.Sprintf("CREATE DATABASE `%s` DEFAULT CHARSET utf8 COLLATE utf8_general_ci;", dbname)
	dbconfig.Dbname = ""
	db := New(dbconfig)
	db.Open(300,100)
	num, err := db.Create(sql)
	dbconfig.Dbname = dbname
	return num, err

}

//打开数据库连接 open a connecttion
//username:password@protocol(address)/dbname?param=value
func (db *Mysql) Open(maxopen int,maxidle int) {
	if db.Client != nil {
		return
	}
	dbs, err := sql.Open("mysql", db.Config.Username+":"+db.Config.Password+"@tcp("+db.Config.Ip+":"+db.Config.Port+")/"+db.Config.Dbname+"?charset=utf8")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	//defer dbs.Close()
	dbs.SetMaxIdleConns(maxidle)
	dbs.SetMaxOpenConns(maxopen)

	err = dbs.Ping()
	if err != nil {
		log.Println("Ping err:"+err.Error())
		log.Fatal(err.Error())
	}

	db.Client = dbs
}

//查询数据库 Query
func (db *Mysql) Select(prestring string, parm ...interface{}) (returnrows []map[string]interface{}, err error) {
	returnrows = []map[string]interface{}{}
	rows, err := db.Client.Query(prestring, parm...)
	if err != nil {
		return
	}

	defer rows.Close()
	// Get column names
	columns, err := rows.Columns()

	if err != nil {
		return nil, err
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		returnrow := map[string]interface{}{}
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			returnrow[columns[i]] = value
			//log.Println(columns[i], ": ", value)

		}
		returnrows = append(returnrows, returnrow)
		//log.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return
}
