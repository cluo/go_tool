/*
 * Created by 一只尼玛 on 2016/8/12.
 * 功能： 网站爬取功能
 *
 */
package spider

import (
	"github.com/op/go-logging"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var Log = logging.MustGetLogger("go_tool_spider")
var format = logging.MustStringFormatter(
	"%{color}%{time:2006-01-02 15:04:05.000} %{longpkg}:%{longfunc} [%{level:.5s}]:%{color:reset} %{message}",
)

// init log record
func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
	logging.SetLevel(logging.INFO, "go_tool_spider")
}

type Spider struct {
	Url    string
	Method string //Get Post
	Header http.Header
	Data   url.Values
	Wait   time.Duration
	Client *http.Client
}

// level name you can refer
var LevelNames = []string{
	"CRITICAL",
	"ERROR",
	"WARNING",
	"NOTICE",
	"INFO",
	"DEBUG",
}

// set log level
func (this *Spider) SetLogLevel(level string) {
	lvl, _ := logging.LogLevel(level)
	logging.SetLevel(lvl, "go_tool_spider")
}

// auto decide which method
func (this *Spider) Go() (body []byte, e error) {
	if strings.ToLower(this.Method) == "post" {
		return this.Post()
	} else {
		return this.Get()
	}

}

// Get method,can take a client
func (this *Spider) Get() (body []byte, e error) {
	// wait but 0 second not
	Wait(this.Wait)

	//debug,can use SetLogLevel to change
	Log.Debug("GET url:" + this.Url)

	//a new request
	request, _ := http.NewRequest("GET", this.Url, nil)

	//clone a header
	request.Header = CloneHeader(this.Header)

	//debug the header
	OutputMaps("---------request header--------", request.Header)

	//start request
	if this.Client == nil {
		// default client
		this.Client = Client
	}
	response, err := this.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	//debug
	OutputMaps("----------response header-----------", response.Header)
	Log.Debugf("Status：%v:%v", response.Status, response.Proto)

	//设置新Cookie
	//Cookieb = MergeCookie(Cookieb, response.Cookies())

	//返回内容 return bytes
	body, e = ioutil.ReadAll(response.Body)

	return
}

// Post附带信息 can take a client
func (this *Spider) Post() (body []byte, e error) {
	Wait(this.Wait)

	Log.Debug("POST url:" + this.Url)

	var request = &http.Request{}

	//post data
	if this.Data != nil {
		pr := ioutil.NopCloser(strings.NewReader(this.Data.Encode()))
		request, _ = http.NewRequest("POST", this.Url, pr)
	} else {
		request, _ = http.NewRequest("POST", this.Url, nil)
	}
	request.Header = CloneHeader(this.Header)

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	OutputMaps("---------request header--------", request.Header)

	if this.Client == nil {
		this.Client = Client
	}
	response, err := this.Client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	OutputMaps("----------response header-----------", response.Header)
	Log.Debugf("Status：%v:%v", response.Status, response.Proto)

	body, e = ioutil.ReadAll(response.Body)

	//设置新Cookie
	//MergeCookie(Cookieb, response.Cookies())

	return
}

// class method
func (this *Spider) NewHeader(ua interface{}, host string, refer interface{}) {
	this.Header = NewHeader(ua, host, refer)
}

// return global log
func (this *Spider) Log() *logging.Logger {
	return Log
}
