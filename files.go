package main

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
)

// joins 2 files together and creates a new one
func comb(file1, file2, new string) {
	if _, err := os.Stat(new); !os.IsNotExist(err) {
		color.Red("Could not create new file: %v", err)
		return
	}
	var chan1 = make(chan []byte)
	var chan2 = make(chan []byte)
	file, err := os.Create(new)
	if err != nil {
		color.Red("error while creating the file: %v", new)
	}
	err = file.Close()
	if err != nil {
		color.Red("error closing file: %v", new)
	}

	file, err = os.OpenFile(new, os.O_RDWR, 0755)
	if err != nil {
		color.Red("gobash> Cat had a problem opening the newly created file: %v\n", new)
	}
	go concurrentReadFile(file1, chan1)
	go concurrentReadFile(file2, chan2)
	fmt.Println("waiting")
	fmt.Println("testing stuffe")
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

// does as it says it reads a file
func readFile(filename string) ([]byte, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0755)
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
		color.Red("gobash> error: %v", err)
		tchan <- []byte("failed to read file")
	}
	tchan <- bytes
	return
}

// does as it says reads bytes from a file till there is no more file left
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
