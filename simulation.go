package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// change input string to usable commands
func command(s string, e error) []string {
	if e != nil {
		log.Println(e)
		return []string{}
	}
	s = strings.TrimSpace(s)
	return strings.Split(s, " ")
}

// shows all files and directory's in a path
func showContent(dir *string) error {
	files, err := ioutil.ReadDir(*dir)
	fmt.Println("gobash> showing files")
	for _, v := range files {
		if v.IsDir() {
			color.Blue("%v", v.Name())
		} else {
			color.Green("%v", v.Name())
		}
	}
	return err
}

// set active working directory
func setDir(s string, dir *string) {
	if len(s) < 1 {
		return
	}
	if strings.Contains(s, "../") || strings.Contains(s, "..") {
		dirUp(s, dir)
	} else {
		dirDown(s, dir)
	}
}

// move up trough the directory's
func dirUp(s string, dir *string) {
	var up int
	if len(s) > 2 {
		up = strings.Count(s, "..")
	} else {
		up = 1
	}
	max := strings.Count(*dir, "/")
	if up > max {
		fmt.Println("gobash> You can't go up that many directories")
		return
	}
	*dir = buildPath(max-up, dir)
}

// move down the directory's
func dirDown(s string, dir *string) {
	if checkExistence(s, dir) {
		*dir += "/" + s
	} else {
		color.Red("%v", fmt.Sprintf("gobash> %v: ./%v %s", "file or directory", s, "not found"))
	}
}

// build the directory path
func buildPath(max int, dir *string) string {
	var tmp string
	t := strings.Split(*dir, "/")
	for k, v := range t {
		if k == max {
			tmp += v
			break
		}
		tmp += v + "/"
	}
	return tmp
}

// checks if a given file exists, file may be removed or changed between checking and returning
func checkExistence(s string, dir *string) bool {
	fileInfo, err := os.Stat(*dir + "/" + s)
	if err != nil && fileInfo.IsDir() {
		return false
	}
	return true
}

// shows information about the commands and how to use them
func showHelp(s string) {
	var commands = setHelpCommands()
	fmt.Println("gobash> showing help")
	switch s {
	case "all":
		color.Green("%v\n", commands["help"])
		color.Green("%v\n", commands["cd"])
		color.Green("%v\n", commands["ls"])
		color.Green("%v\n", commands["exec"])
		color.Green("%v\n", commands["exit"])
	default:
		color.Red("help command unrecognized")
	}
}
