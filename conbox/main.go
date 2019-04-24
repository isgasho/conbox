package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/udhos/conbox/applets/cat"
	"github.com/udhos/conbox/applets/echo"
	"github.com/udhos/conbox/applets/ls"
	"github.com/udhos/conbox/applets/pwd"
	"github.com/udhos/conbox/applets/rm"
	"github.com/udhos/conbox/common"
)

func main() {

	appletTable := loadApplets()

	// 1. try basename
	appletName := filepath.Base(os.Args[0])
	if applet, found := appletTable[appletName]; found {
		run(applet, os.Args[1:])
		return
	}

	if appletName != "conbox" {
		common.ShowVersion()
		fmt.Printf("conbox: basename: applet '%s' not found\n", appletName)
		usage(appletTable)
		os.Exit(1)
	}

	// 2. try arg 1
	if len(os.Args) > 1 {
		arg := os.Args[1]
		switch arg {
		case "-h":
			common.ShowVersion()
			usage(appletTable)
			return
		case "-l":
			listApplets(appletTable, "\n")
			return
		}
		appletName = arg
		if applet, found := appletTable[appletName]; found {
			run(applet, os.Args[2:])
			return
		}
		common.ShowVersion()
		fmt.Printf("conbox: arg 1: applet '%s' not found\n", appletName)
		usage(appletTable)
		os.Exit(2)
	}

	common.ShowVersion()
	usage(appletTable)
	os.Exit(3)
}

func usage(tab map[string]appletFunc) {
	fmt.Println("usage: conbox APPLET [ARG]... : run APPLET")
	fmt.Println("       conbox -h              : show command-line help")
	fmt.Println("       conbox -l              : list applets")
	fmt.Println()
	fmt.Println("conbox: registered applets:")
	listApplets(tab, " ")
	fmt.Println()
}

func listApplets(tab map[string]appletFunc, sep string) {
	var list []string
	for n := range tab {
		list = append(list, n)
	}
	sort.Strings(list)
	for _, n := range list {
		fmt.Printf("%s%s", n, sep)
	}
}

func run(applet appletFunc, args []string) {
	exitCode := applet(args)
	if exitCode != 0 {
		os.Exit(exitCode)
	}
}

type appletFunc func(args []string) int

func loadApplets() map[string]appletFunc {
	tab := map[string]appletFunc{
		"cat":  cat.Run,
		"echo": echo.Run,
		"ls":   ls.Run,
		"pwd":  pwd.Run,
		"rm":   rm.Run,
	}
	return tab
}
