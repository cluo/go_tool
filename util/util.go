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

//def devidelist(files, num=0):
//    filestype = type(files)
//    if not filestype == type([]):
//        raise Exception("文件切分只能是列表")
//    length = len(files)
//    split = {}
//    if length <= 0:
//        return split
//    if num >= length:
//        raise Exception("文件列表切分过小")
//    process = length // num
//    for i in range(num):
//        split[i] = (files[i * process:(i + 1) * process])
//    remain = files[num * process:]
//    for i in range(len(remain)):
//        split[i % num].append(remain[i])
//    return split
func DevideStringList(files []string, num int) (map[int][]string, error) {
	length := len(files)
	split := map[int][]string{}
	if length <= 0 {
		return split, errors.New("num must not negtive")
	}
	if num >= length {
		return split, errors.New("num must not bigger than the list length")
	}
	process := length / num
	for i := 0; i < num; i++ {
		split[i] = (files[i*process : (i+1)*process])
	}
	remain := files[num*process:]
	for i := 0; i < len(remain); i++ {
		split[i%num] = append(split[i%num], remain[i])
	}
	return split, nil
}
