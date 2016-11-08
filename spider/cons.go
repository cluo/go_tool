//常量包
package spider

const (
	//暂停时间 default wait time
	WaitTime = 5
)

var (
	//浏览器头部 default header ua
	FoxfireLinux = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:46.0) Gecko/20100101 Firefox/46.0"
	SpiderHeader = map[string][]string{
		"User-Agent": {
			FoxfireLinux,
		},
	}
)

// usually a header has ua,host and refer
func NewHeader(ua interface{}, host string, refer interface{}) map[string][]string {
	if ua == nil {
		ua = FoxfireLinux
	}
	if refer == nil {
		h := map[string][]string{
			"User-Agent": {
				ua.(string),
			},
			"Host": {
				host,
			},
		}
		return h
	}
	h := map[string][]string{
		"User-Agent": {
			ua.(string),
		},
		"Host": {
			host,
		},
		"Referer": {
			refer.(string),
		},
	}
	return h
}
