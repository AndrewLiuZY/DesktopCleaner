package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	common "../util"
)

const dbType string = "sqlite3"

const dbName string = "E:/Github/DesktopCleaner/src/db/dc.db"

func getDB() *sql.DB {
	db, err := sql.Open(dbType, common.Config()["dbname"])
	checkErr(err)
	return db
}

//AddOutSide 增加一个不分类的扩展名
func AddOutSide(ex string) {
	db := getDB()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO OUTSIDE(EXTENSION) values(?)")
	checkErr(err)
	res, err := stmt.Exec(ex)
	checkErr(err)
	_, err = res.RowsAffected()
	checkErr(err)
}

//RemoveOutSide 删除一个不分类的扩展名
func RemoveOutSide(ex string) {
	db := getDB()
	defer db.Close()

	stmt, err := db.Prepare("delete from OUTSIDE where EXTENSION=?")
	checkErr(err)
	res, err := stmt.Exec(ex)
	checkErr(err)
	_, err = res.RowsAffected()
	checkErr(err)

}

//BindExtension 绑定扩展名和文件夹
func BindExtension(ex, dir string) {
	db := getDB()
	defer db.Close()
	//删除旧的绑定
	stmt, err := db.Prepare("delete from DIR where EXTENSION=?")
	checkErr(err)
	res, err := stmt.Exec(ex)
	checkErr(err)
	_, err = res.RowsAffected()
	checkErr(err)

	stmt, err = db.Prepare("INSERT INTO DIR(DIRNAME,EXTENSION) values(?,?)")
	checkErr(err)
	res, err = stmt.Exec(dir, ex)
	checkErr(err)
	_, err = res.RowsAffected()
	checkErr(err)
}

//RemoveBind 根据文件名删除绑定
func RemoveBind(ex string) {
	db := getDB()
	defer db.Close()

	stmt, err := db.Prepare("delete from DIR where EXTENSION=?")
	checkErr(err)
	res, err := stmt.Exec(ex)
	checkErr(err)
	_, err = res.RowsAffected()
	checkErr(err)
}

//DisplayDIR 显示所有绑定信息
func DisplayDIR() {
	db := getDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM DIR")
	checkErr(err)

	for rows.Next() {
		var DIR string
		var EXTENSION string
		err = rows.Scan(&DIR, &EXTENSION)
		checkErr(err)
		fmt.Println(DIR, EXTENSION)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
