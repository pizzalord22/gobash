package main

import (
	"fmt"
	"github.com/fatih/color"
)

// shows information about the commands and how to use them
func showHelp(s string) {
	var commands = setHelpCommands()
	fmt.Println("gobash> showing help")
	switch s {
	case "all", "help":
		color.Green("%v\n", commands["help"])
		color.Green("%v\n", commands["cd"])
		color.Green("%v\n", commands["ls"])
		color.Green("%v\n", commands["exec"])
		color.Green("%v\n", commands["exit"])
	case "cd":
		color.Green("%v\n", commands["cd"])
	case "ls":
		color.Green("%v\n", commands["ls"])
	case "exec":
		color.Green("%v\n", commands["exec"])
	case "exit":
		color.Green("%v\n", commands["exit"])
	default:
		color.Red("help command unrecognized")
	}
}

// set the text to show when the help command is ran
func setHelpCommands() map[string]string {
	var hc = make(map[string]string)
	hc["help"] = "help\n----------\nshows this help window\n"
	hc["cd"] = "cd\n----------\nuse cd file/dir to go to a directory, use ../ to go up a file\n"
	hc["ls"] = "ls\n----------\nlist all files and directories in current path\n"
	hc["exec"] = "ls\n----------\nthis allows you to execute commands on the underlying os that gobash runs on.\nMeaning you can use native linux/mac and windows command and get their output\n"
	hc["exit"] = "exit\n----------\nNicely exits the simulated\n"
	return hc
}
