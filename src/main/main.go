package main

import (
	"../db"
)

func main() {
	// user, err := user.Current()
	// util.Check(err)
	// desktopPath := user.HomeDir + "/DeskTop/test/test"
	// clean.Clean(desktopPath)
	db.BindExtension("png","Images")
	db.DisplayDIR()
}
