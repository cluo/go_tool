/*
 * Created by 一只尼玛 on 2016/8/12.
 * 功能： 杂类
 *
 */
package util

import (
	"strconv"
	"strings"
	"time"
)

//string to int
func SI(s string) (i int, e error) {
	i, e = strconv.Atoi(s)
	return
}

//int to string
func IS(i int) string {
	return strconv.Itoa(i)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

// sleep
func Sleep(waittime int) {
	time.Sleep(time.Duration(waittime) * time.Second)
}

//time
func Second(times int) time.Duration {
	return time.Duration(times) * time.Second
}


func TodayString(level int) string {
	formats := "20060102-15:04:05"
	switch level {
	case 1:
		formats = "2006"
	case 2:
		formats = "200601"
	case 3:
		formats = "20060102"
	case 4:
		formats = "20060102-15"
	case 5:
		formats = "20060102-15:04"
	default:

	}
	return time.Now().Format(formats)
}
