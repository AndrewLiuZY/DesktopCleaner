package util

import (
	"fmt"
	"runtime"
)

//Check check and print the errors;
//err==nil return true ;
//err!=nil print log and return false
func Check(err error) bool{
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(file,line, err.Error())
		return false
	}
	return true
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
