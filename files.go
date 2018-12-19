package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
)

// TODO: Create function for writing to files


// does as it says it reads a file
func readFile(filename string) ([]byte, error) {
	file, err := os.OpenFile(filename,  0x00000, 0755)
	defer file.Close()
	if err != nil {
		return []byte{}, nil
	}
	return readBytes(file)
}

// this function should be used to read multiple files concurrently
func concurrentReadFile(filename string, tchan chan []byte) {
	bytes, err := readFile(filename)
	if err != nil {
		color.Red("gobash> error: %v", errors.Wrap(err, fmt.Sprintf("error while readind the file %s: ", filename)))
		tchan <- []byte(fmt.Sprintf("Failed to read file: %s", filename))
	}
	tchan <- bytes
	return
}

// does as it says reads bytes from a file till there is no more file left
// TODO read files with a buffer size so not everything is kept in memory
func readBytes(file *os.File) (r []byte, e error) {
	buffer := make([]byte, bufferSize)
	for {
		_, e = file.Read(buffer)
		if e != nil {
			if e != io.EOF {
				return []byte{}, e
			}
			break
		}
		for _, v := range buffer {
			r = append(r, v)
		}
	}
	return r, nil
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

// checks if a given file exists, file may be removed or changed between checking and returning
func checkExistence(s string, dir *string) bool {
	fileInfo, err := os.Stat(*dir + "/" + s)
	if err != nil && fileInfo.IsDir() {
		return false
	}
	return true
}
