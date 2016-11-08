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
