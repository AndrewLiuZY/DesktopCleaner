package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
	"strings"

	"../util"
)

func main() {
	user, err := user.Current()
	util.Check(err)
	basePath := user.HomeDir + "/DeskTop"
	files, err := ioutil.ReadDir(basePath)
	util.Check(err)
	for _, file := range files {
		if i := strings.LastIndex(file.Name(), "."); i >= 1 {
			name := file.Name()[i+1:]
			if !checkFileValid(name) { //除去桌面快捷方式
				continue
			}
			dirPath := basePath + "/" + name
			os.Mkdir(dirPath, os.ModeDir)
			srcPath := basePath + "/" + file.Name()
			dstPath := dirPath + "/" + file.Name()
			copy(srcPath, dstPath)
		}
	}
}

func checkFileValid(t string) bool {
	return t != "lnk"
}

func copy(src, dst string) {
	reader, err := os.Open(src)
	util.Check(err)
	defer reader.Close()
	//输出结果
	printResult := func(i int) {
		fmt.Println(src+"  ->  "+dst, " "+strconv.FormatFloat((float64(i)/1024), 'f', 3, 64)+"kb")
	}
	//小文件(<10k)直接读写全部
	if fileinfo, _ := os.Stat(src); fileinfo.Size() < 10*1024 {
		data, err := ioutil.ReadAll(reader)
		util.Check(err)
		err = ioutil.WriteFile(dst, data, os.ModeAppend)
		if util.Check(err) {
			printResult(len(data))
		}
	} else { //超过10k的文件直接用io.Copy
		writer, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
		util.Check(err)
		defer writer.Close()
		length, err := io.Copy(writer, reader)
		if util.Check(err) {
			printResult(int(length))
		}
	}
}

func getDirName(fileName string) {

}
