/*
 * Created by 一只尼玛 on 2016/8/12.
 * 功能： 文件帮助功能
 *
 */
package util

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

//获取调用者的当前文件DIR
//Get the caller now directory
func CurDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Dir(filename)
}

//将字节数组保存到文件中去
//Save bytes into file
func SaveToFile(filepath string, content []byte) error {
	//全部权限写入文件
	err := ioutil.WriteFile(filepath, content, 0777)
	return err
}

// read bytes from file
func ReadfromFile(filepath string) ([]byte, error) {
	return ioutil.ReadFile(filepath)
}

//根据传入文件夹名字递归新建文件夹
//Create dir by recursion
func MakeDir(filedir string) error {
	return os.MkdirAll(filedir, 0777)
}

//根据传入文件名，递归创建文件夹
// ./dir/filename  /home/dir/filename
//Create dir by the filename
func MakeDirByFile(filepath string) error {
	temp := strings.Split(filepath, "/")
	if len(temp) <= 2 {
		return errors.New("please input complete file name like ./dir/filename or /home/dir/filename")
	}
	dirpath := strings.Join(temp[0:len(temp)-1], "/")
	return MakeDir(dirpath)
}

func FileExist(filename string) bool {
	fi, err := os.Stat(filename)
	if err != nil {
		return false
	}
	if fi.IsDir() {
		return false
	} else {
		return true
	}
}

func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {

		if fi.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix)
	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}
