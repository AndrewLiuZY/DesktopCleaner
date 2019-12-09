package main

import (
	"os/user"

	"../clean"
	"../util"
)

func main() {
	user, err := user.Current()
	util.Check(err)
	desktopPath := user.HomeDir + "/DeskTop/test/test"
	clean.Clean(desktopPath)
}
