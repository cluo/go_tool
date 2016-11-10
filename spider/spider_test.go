package spider

import (
	"testing"
	"fmt"
)

func TestSpider(t *testing.T) {
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

	// global log record
	//SetLogLevel("DEBUg")
	SetLogLevel("error")

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

	// if filesize small than 500KB
	err = TooSortSizes(body, 500)

	Log().Error(err.Error())
}
