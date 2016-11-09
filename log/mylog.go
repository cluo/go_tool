package mylog

// simple log,new rotate func
import (
	"github.com/hunterhug/go_tool/util"
	"log"
	"os"
)

// my log
type Log struct {
	LogFilename string
	//FileSize    int64
	Logger   *log.Logger
	LogLevel string // error < info < debug < all
}

func NewLog(logfilename string) Log {
	temp := Log{}
	temp.LogFilename = logfilename
	temp.LogLevel = "ALL"
	//temp.FileSize = filesize
	file, err := os.OpenFile(logfilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("fail to create test.log file!")
		panic("no log")
	}
	//fileinfo, _ := file.Stat()
	////  if filesize small than ?M
	//if fileinfo.Size() > filesize/(1024*1024) {
	//	//Todo
	//	//??
	//}
	logger := log.New(file, "[I]:", log.LstdFlags|log.Lshortfile)
	// print to conslose
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	temp.Logger = logger
	return temp
}

func (this *Log) Errorf(format string, v ...interface{}) {
	temp := util.ToLower(this.LogLevel)
	if temp == "error" || temp == "debug" || temp == "info" || temp == "all" {
		log.Printf("[E]:"+format, v...)
		this.Logger.SetPrefix("[E]:")
		this.Logger.Printf(format, v...)
	}
}

func (this *Log) Debugf(format string, v ...interface{}) {
	temp := util.ToLower(this.LogLevel)
	if temp == "debug" || temp == "all" {
		log.Printf("[D]:"+format, v...)
		this.Logger.SetPrefix("[D]:")
		this.Logger.Printf(format, v...)
	}
}

func (this *Log) Error(v ...interface{}) {
	temp := util.ToLower(this.LogLevel)
	if temp == "error" || temp == "debug" || temp == "info" || temp == "all" {
		v[0] = "[E]:" + v[0].(string)
		log.Println(v...)
		this.Logger.SetPrefix("[E]:")
		this.Logger.Println(v...)
	}
}

func (this *Log) Debug(v ...interface{}) {
	temp := util.ToLower(this.LogLevel)
	if temp == "debug" || temp == "all" {
		v[0] = "[D]:" + v[0].(string)
		log.Println(v...)
		this.Logger.SetPrefix("[D]:")
		this.Logger.Println(v...)
	}
}

func (this *Log) Println(v ...interface{}) {
	temp := util.ToLower(this.LogLevel)
	if temp == "debug" || temp == "info" || temp == "all" {
		v[0] = "[I]:" + v[0].(string)
		log.Println(v...)
		this.Logger.SetPrefix("[I]:")
		this.Logger.Println(v...)
	}
}

func (this *Log) Printf(format string, v ...interface{}) {
	temp := util.ToLower(this.LogLevel)
	if temp == "debug" || temp == "info" || temp == "all" {
		log.Printf("[I]:"+format, v...)
		this.Logger.SetPrefix("[I]:")
		this.Logger.Printf(format, v...)
	}
}

func (this *Log) SetLevel(level string) {
	temp := util.ToLower(level)
	if temp == "debug" || temp == "info" || temp == "all" || temp == "info" || temp == "error" {
		this.LogLevel = level
	} else {
		panic("Level error")
	}
}
