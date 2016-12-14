/*
 * Created by 一只尼玛 on 2016/8/12.
 * 功能： 杂类
 *
 */
package util

import (
	"errors"
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

// get secord times
//172606056
func GetSecordTimes() int64 {
	return time.Now().Unix()
}

//201611112113
func GetSecord2DateTimes(secord int64) string {
	tm := time.Unix(secord, 0)
	return tm.Format("20060102150405")

}

func GetDateTimes2Secord(datestring string) int64 {
	tm2, _ := time.Parse("20060102150405", datestring)
	return tm2.Unix()

}
func TodayString(level int) string {
	formats := "20060102150405"
	switch level {
	case 1:
		formats = "2006"
	case 2:
		formats = "200601"
	case 3:
		formats = "20060102"
	case 4:
		formats = "2006010215"
	case 5:
		formats = "200601021504"
	default:

	}
	return time.Now().Format(formats)
}

// change by python
func DevideStringList(files []string, num int) (map[int][]string, error) {
	length := len(files)
	split := map[int][]string{}
	if num <= 0 {
		return split, errors.New("num must not negtive")
	}
	if num >= length {
		return split, errors.New("num must not bigger than the list length")
	}
	process := length / num
	for i := 0; i < num; i++ {
		// slice inside has a refer, so must do this append
		//split[i]=files[i*process : (i+1)*process] wrong!
		split[i] = append(split[i], files[i*process:(i+1)*process]...)
	}
	remain := files[num*process:]
	for i := 0; i < len(remain); i++ {
		split[i%num] = append(split[i%num], remain[i])
	}
	return split, nil
}

//字符串是否在字符串数组中
func InArray(sa []string, a string) bool {
	for _, v := range sa {
		if a == v {
			return true
		}
	}
	return false
}

func StringsToJson(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}

	return jsons
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}
