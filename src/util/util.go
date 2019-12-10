package util

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
)

//Check check and print the errors;
//err==nil return true ;
//err!=nil print log and return false
func Check(err error) bool {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(file, line, err.Error())
		return false
	}
	return true
}

//PanicCheck err!=null => panic
func PanicCheck(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		panic(fmt.Sprintln(file, line, err.Error()))
	}
}

//Contain Contain
func Contain(s []string, i string) bool {
	for _, item := range s {
		if item == i {
			return true
		}
	}
	return false
}

//ForEach foreach
func ForEach(s []string, f func(index int, str string)) {
	for i, item := range s {
		f(i, item)
	}
}

//FileExist 根据文件路径检查文件是否存在
func FileExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsExist(err)
}

//Config 获取配置信息
func Config() map[string]string {
	file, err := os.OpenFile("../Data/config.json", os.O_RDONLY, 0666)
	PanicCheck(err)
	defer file.Close()
	deCoder := json.NewDecoder(file)
	config := make(map[string]string)
	deCoder.Decode(&config)
	return config
}
