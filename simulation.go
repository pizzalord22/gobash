package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
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

// set active working directory
func setDir(s string, dir *string) {
	if len(s) < 1 {
		return
	}
	if strings.Contains(s, "/") {
		subPaths := strings.Split(s, "/")
		fmt.Println(subPaths)
		for _, v := range subPaths {
			fmt.Println(*dir)
			if v == ".." {
				dirUp(dir)
			} else {
				dirDown(v, dir)
			}
		}
	} else {
		if s == ".." {
			dirUp(dir)
		} else {
			dirDown(s, dir)
		}
	}
}

// move up trough the directory's
func dirUp(dir *string) {
	var tmp string
	t := strings.Split(*dir, "/")
	pl := strings.Count(*dir, "/")
	if len(t) == 1 {
		fmt.Println("you can't go higher")
		return
	}
	for k, v := range t {
		if k == pl-1 {
			tmp += v
			break
		}
		tmp += v + "/"
	}
	*dir = tmp
}

// move down the dirtodoectory's
func dirDown(s string, dir *string) {
	if checkExistence(s, dir) {
		*dir += "/" + s
	} else {
		color.Red("%v", fmt.Sprintf("gobash> %v: ./%v %s", "file or directory", s, "not found"))
	}
}
