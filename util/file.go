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
	if fi.IsDir(){
		return false
	}else {
		return true
	}
}
