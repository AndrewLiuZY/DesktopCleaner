package clean

import (
	"encoding/gob"
	"os"

	"../util"
)

type record map[string][]string

func newRecord() record {
	return make(map[string][]string)
}

func (r record) addRecord(dir, file string) {
	if _, ok := r[dir]; !ok {
		r[dir] = make([]string, 0)
	}
	r[dir] = append(r[dir], file)
}

func (r record) save() {
	path, _ := util.GetCurrentPath()
	path += "data"
	os.Mkdir(path, os.ModeDir)
	file, err := os.OpenFile(path+"/cache.gob", os.O_RDONLY|os.O_CREATE, 0666)
	defer file.Close()
	util.PanicCheck(err)
	enc := gob.NewEncoder(file)
	err = enc.Encode(r)
	util.Check(err)
}
