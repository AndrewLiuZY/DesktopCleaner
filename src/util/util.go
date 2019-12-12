package util

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
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
func Contain(s []interface{}, i string) bool {
	for _, item := range s {
		str, _ := item.(string)
		if str == i {
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
	return !os.IsNotExist(err)
}

//GetPathFromConfig 从配置文件中获取path
func GetPathFromConfig() string {
	path, ok := Config()["path"].(string)
	if !ok {
		panic("配置文件中\"path\"错误或不存在")
	}
	return path
}

//Config 获取配置信息
func Config() map[string]interface{} {
	path, _ := GetCurrentPath()
	file, err := os.OpenFile(path+"\\config.json", os.O_RDONLY, 0666)
	PanicCheck(err)
	defer file.Close()
	deCoder := json.NewDecoder(file)
	config := make(map[string]interface{})
	deCoder.Decode(&config)
	return config
}

//GetCurrentPath 获取当前可执行文件路径
func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	return string(path[0 : i+1]), nil
}
