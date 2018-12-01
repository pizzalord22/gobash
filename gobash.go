package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const bufferSize = 100

// main loop here we initialize our reader
func main() {
	reader := bufio.NewReader(os.Stdin)
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	dir := strings.Replace(filepath.Dir(ex), "\\", "/", -1)
	fmt.Printf("%v>", dir)
	osLoop(reader, &dir)
}

// this stimulates an os loop
func osLoop(reader *bufio.Reader, dir *string) {
	var err error
	for {
		err = nil
		r := command(reader.ReadString('\n'))
		switch r[0] {
		case "exit":
			return
		case "ls":
			err = showContent(dir)
		case "cd":
			if len(r) > 1 {
				setDir(r[1], dir)
			} else {
				showHelp("cd")
			}
		case "help", "h", "-help", "-h":
			showHelp("all")
		case "exec":
			if len(r) > 1 {
				execute(r[1:])
			} else {
				showHelp("exec")
			}
		case "comb":
			if len(r) > 2 {
				comb(r[1], r[2], r[3])
			} else {
				showHelp("comb")
			}
		case "webdog":
			execute([]string{"start", "chrome.exe", "mydogchase.com"})
		default:
			fmt.Printf("gobash> Type \"help\" to see a list of available commands\n")
		}
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%v>", *dir)
	}
}
