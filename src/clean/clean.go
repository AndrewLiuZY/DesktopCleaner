package clean

import (
	"../util"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//Clean clean the dir
func Clean(dirName string) {
	basePath := dirName
	files, err := ioutil.ReadDir(basePath)
	util.Check(err)
	for _, file := range files {
		if i := strings.LastIndex(file.Name(), "."); i >= 1 {
			name := file.Name()[i+1:]
			if !fileValid(name) { //除去桌面快捷方式
				continue
			}
			dirPath := basePath + "/" + name
			os.Mkdir(dirPath, os.ModeDir)
			srcPath := basePath + "/" + file.Name()
			dstPath := dirPath + "/" + file.Name()
			copy(srcPath, dstPath)
			confirm(srcPath, dstPath)
		}
	}
}

//校验
func confirm(srcPath, dstPath string) {
	if sameFile(srcPath, dstPath) {
		//复制完成后删除源文件
		err := os.Remove(srcPath)
		util.Check(err)
	} else {
		err := os.Remove(dstPath)
		util.Check(err)
		fmt.Println("文件:", srcPath, "移动失败！")
	}
}

//对比俩文件的size
func sameFile(src, dst string) bool {
	srcFile, err := os.Stat(src)
	util.Check(err)
	dstFile, err := os.Stat(dst)
	util.Check(err)
	return srcFile.Size() == dstFile.Size()
}

//复制文件
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

//根据扩展名获取文件夹名 TODO:使用sqlite存储
func getDirName(fileName string) {
	image := new(fileDir)
	image.dirName = "Images"
	image.extensions = []string{"jpg", "png", "gif", "bmp", "tif", "pcx", "tga", "exif", "fpx", "svg", "psd", "cdr", "pcd", "dxf", "ufo", "eps"}
}

//检查文件扩展名是否有效 TODO:使用sqlite存储
func fileValid(t string) bool {
	return t != "lnk"
}

type fileDir struct {
	dirName    string
	extensions []string
}
