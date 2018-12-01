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
