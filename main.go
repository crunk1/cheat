package main

import (
	"fmt"
	"os"
)

const USAGE = `Usage:
  cheat <cheatsheet>
  cheat -d <cheatsheet>
  cheat -e <cheatsheet>
  cheat -h
  cheat -l

Cheatsheet manager. Displays, edits, and deletes cheatsheets.

Options:
  -d Deletes a cheatsheet.
  -e Edits a cheatsheet. If the cheatsheet does not exist, it is created.
  -h Displays this help message.
  -l Lists cheatsheets.

Examples:
  % cheat -e grep
  % cheat grep`

var cmdsMap = map[string]func(args []string){
	"-d":   DeleteCmd,
	"-e":   EditCmd,
	"-h":   HelpCmd,
	"-l":   ListCmd,
	"show": ShowCmd,
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		badUsageExit("Too few arguments.")
	}

	if args[0][0] == '-' {
		cmd, ok := cmdsMap[args[0]]
		if !ok {
			badUsageExit(fmt.Sprintf("Unknown flag: %s\n", args[0]))
		}
		cmd(args)
	} else if len(args) != 1 {
		badUsageExit("Too many arguments.")
	} else {
		cmdsMap["show"](args)
	}
}

func badUsageExit(msg string) {
	if msg != "" {
		fmt.Println(msg)
	}
	fmt.Println(USAGE)
	os.Exit(1)
}

func errorExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
