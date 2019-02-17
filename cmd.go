package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"sync"

	"github.com/mitchellh/go-homedir"
)

var (
	CHEATDIR   = ""
	cheatDirMx = sync.Mutex{}
)

func getCheatDir() string {
	cheatDirMx.Lock()
	defer cheatDirMx.Unlock()
	if CHEATDIR == "" {
		homeD, err := homedir.Dir()
		if err != nil {
			errorExit(err)
		}
		CHEATDIR = path.Join(homeD, ".cheat")
		if err = os.MkdirAll(CHEATDIR, 0700); err != nil {
			errorExit(err)
		}
	}
	return CHEATDIR
}

func DeleteCmd(args []string) {
	if len(args) != 2 {
		badUsageExit("-e takes exactly one argument")
	}
	cheatsheet := args[1]
	err := os.Remove(path.Join(getCheatDir(), cheatsheet))
	if err != nil {
		errorExit(err)
	}
}

func EditCmd(args []string) {
	if len(args) != 2 {
		badUsageExit("-e takes exactly one argument")
	}
	cheatsheet := args[1]
	p := path.Join(getCheatDir(), cheatsheet)

	cmd := exec.Command("touch", p)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		errorExit(err)
	}

	cmd = exec.Command("vim", p)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		errorExit(err)
	}
}

func HelpCmd(args []string) {
	if len(args) > 1 {
		badUsageExit("-h takes no arguments")
	}
	fmt.Println(USAGE)
}

func ListCmd(args []string) {
	if len(args) > 1 {
		badUsageExit("-l takes no arguments")
	}
	files, err := ioutil.ReadDir(getCheatDir())
	if err != nil {
		errorExit(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func ShowCmd(args []string) {
	cheatsheet := args[0]
	cmd := exec.Command("less", path.Join(getCheatDir(), cheatsheet))
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		errorExit(err)
	}
}
