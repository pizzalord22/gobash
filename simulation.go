package main

import (
	"fmt"
	"github.com/fatih/color"
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

// joins 2 files together and creates a new one
// TODO: rework to cat http://www.linfo.org/cat.html
func cat(file1, file2, new string) {
	if _, err := os.Stat(new); !os.IsNotExist(err) {
		color.Red("File already exists: %v", err)
		return
	}
	var chan1 = make(chan []byte)
	var chan2 = make(chan []byte)
	file, err := os.Create(new)
	if err != nil {
		color.Red("Error while creating file: %v", new)
	}
	err = file.Close()
	if err != nil {
		color.Red("There was an err while closing the file: %v", new)
	}

	file, err = os.OpenFile(new, os.O_RDWR, 0755)
	if err != nil {
		color.Red("gobash> Comb had a problem opening the newly created file: %v\n", new)
	}
	go concurrentReadFile(file1, chan1)
	go concurrentReadFile(file2, chan2)
	fmt.Println("waiting")
	_, err = file.Write(<-chan1)
	if err != nil {
		color.Red("gobash> There was a problem writing to %v", file.Name())
	}
	_, err = file.Write(<-chan2)
	if err != nil {
		color.Red("gobash> There was a problem writing to %v", file.Name())
	}
	return
}