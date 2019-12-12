package clean

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"../util"
)

//Recover 从缓存文件中恢复
func Recover(basePath string) {
	record, ok := getCache()
	if !ok {
		util.PanicCheck(errors.New("缓存文件不存在或错误"))
	}
	for dir, files := range record {
		for _, file := range files {
			dst := basePath + "/" + path.Base(file)
			copyFile(file, dst)
			confirm(file, dst)
		}
		//如果文件夹里没有文件就删除文件夹
		files, err := ioutil.ReadDir(dir)
		util.PanicCheck(err)
		if len(files) <= 0 {
			err = os.Remove(dir)
			util.Check(err)
		}
	}
	//还原成功删除缓存文件
	rmCache()
	fmt.Println("还原成功！！\nrecover success!!")
}

func getCache() (map[string][]string, bool) {
	rec := make(map[string][]string)
	basePath, err := util.GetCurrentPath()
	util.PanicCheck(err)
	file, err := os.Open(basePath + "/data/cache.gob")
	defer file.Close()
	if err != nil {
		return nil, false
	}
	dec := gob.NewDecoder(file)
	err = dec.Decode(&rec)
	util.PanicCheck(err)
	return rec, true
}

func rmCache() {
	basePath, err := util.GetCurrentPath()
	util.PanicCheck(err)
	err = os.Remove(basePath + "/data/cache.gob")
	util.Check(err)
}
