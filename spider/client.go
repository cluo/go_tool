/*
 * Created by 一只尼玛 on 2016/8/12.
 * 功能： 网络COOKIE功能
 *
 */
package spider

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

//cookie record
func NewJar() *cookiejar.Jar {
	cookieJar, _ := cookiejar.New(nil)
	return cookieJar
}

var (
	//default client to ask get or post
	Client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			Log.Debugf("-----------Redirect:%v------------", req.URL)
			return nil
		},
		Jar: NewJar(),
	}
	//每次访问携带的cookie not use
	Cookieb = []*http.Cookie{} //map[string][]string
)


// a proxy client
func NewProxyClient(proxystring string) (*http.Client,error){
	proxy, err := url.Parse(proxystring)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		// allow redirect
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			Log.Debugf("-----------Redirect:%v------------", req.URL)
			return nil
		},
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
		Jar: NewJar(),
	}
	return client,nil
}


// a client
func NewClient() (*http.Client,error){
	client := &http.Client{
		// allow redirect
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			Log.Debugf("-----------Redirect:%v------------", req.URL)
			return nil
		},
		Jar: NewJar(),
	}
	return client,nil
}

//merge Cookie，后来的覆盖前来的
func MergeCookie(before []*http.Cookie, after []*http.Cookie) []*http.Cookie {
	cs := make(map[string]*http.Cookie)

	for _, b := range before {
		cs[b.Name] = b
	}

	for _, a := range after {
		if a.Value != "" {
			cs[a.Name] = a
		}
	}

	res := make([]*http.Cookie, 0, len(cs))

	for _, q := range cs {
		res = append(res, q)

	}

	return res

}

// clone a header
func CloneHeader(h map[string][]string) map[string][]string {
	if h == nil {
		h = SpiderHeader
	}
	return CopyM(h)
}
