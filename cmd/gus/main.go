package main

import (
	"github.com/mcandre/gus"

	"flag"
	"fmt"
	"os"
)

var flagTopLevel = flag.String("top", "", "Specify git top-level project path. Default: <current working directory>")
var flagList = flag.Bool("list", false, "List any submodules")
var flagInit = flag.Bool("init", false, "Initialize git repository and submodules")
var flagAdd = flag.String("add", "", "Add the specified submodule URL. Requires -path")
var flagPath = flag.String("path", "", "Target path")
var flagBranch = flag.String("branch", "", "Specify submodule branch")
var flagRemove = flag.String("remove", "", "Remove submodule by URL")
var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

func main() {
	flag.Parse()

	var top string

	if *flagTopLevel != "" {
		top = *flagTopLevel
	} else {
		tp, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		top = tp
	}

	switch {
	case *flagInit:
		if err := gus.Init(top); err != nil {
			panic(err)
		}
	case *flagAdd != "" && *flagPath != "":
		if err := gus.AddSubmodule(top, *flagAdd, *flagPath, *flagBranch); err != nil {
			panic(err)
		}
	case *flagRemove != "":
		if err := gus.RemoveSubmodule(top, *flagRemove); err != nil {
			panic(err)
		}
	case *flagList:
		submodules, err := gus.GetSubmodules(top)

		if err != nil {
			panic(err)
		}

		for url, pth := range submodules {
			fmt.Printf("%v %v\n", url, pth)
		}
	case *flagVersion:
		fmt.Println(gus.Version)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
