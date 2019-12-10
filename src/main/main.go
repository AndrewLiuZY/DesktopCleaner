package main

import (
	"../clean"
	"../util"
)

func main() {
	cleanPath := util.GetPathFormConfig()
	clean.Clean(cleanPath)
}
