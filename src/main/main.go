package main

import (
	"flag"
	"fmt"
	"os"

	"../clean"
	"../util"
)

var (
	h       bool
	r       bool
)

func init() {
	flag.BoolVar(&h, "help", false, "show this help")
	flag.BoolVar(&r, "recover", false, "recover files")

	flag.Usage = usage
}

func main() {
	flag.Parse()
	if h  {
		flag.Usage()
	} else if r  {
		cleanPath := util.GetPathFromConfig()
		clean.Recover(cleanPath)
	} else {
		cleanPath := util.GetPathFromConfig()
		clean.Clean(cleanPath)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `DesktopCleaner version: 0.1.0
Usage: clean [-h help] [-r recover] 

Options:
`)
	flag.PrintDefaults()
}
