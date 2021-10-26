package main

import (
    "fmt"
    "github.com/fatih/color"
    "github.com/pkg/errors"
    "io"
    "io/ioutil"
    "log"
    "os"
)

// TODO: Create function for writing strings to files

// TODO: change the read functions to read a 100 bytes at a time so they are not as memory intensife
// does as it says it reads a file
func readFile(filename string) ([]byte, error) {
    file, err := os.OpenFile(filename, 0x00000, 0755)
    defer file.Close()
    if err != nil {
        return []byte{}, nil
    }
    return readBytes(file)
}

// this function should be used to read multiple files concurrently
func concurrentReadFile(filename string, readChan chan []byte) {
    bytes, err := readFile(filename)
    if err != nil {
        color.Red("gobash> error: %v", errors.Wrap(err, fmt.Sprintf("error while readind the file %s: ", filename)))
        readChan <- []byte(fmt.Sprintf("Failed to read file: %s", filename))
    }
    readChan <- bytes
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

// used to write bytes to a file
// return true if the write as successful
// and returns false if it was not
// it logs the errors if there are any
func writeFile(filename string, text []byte) bool {
    file, err := os.OpenFile(filename, os.O_WRONLY, 0777)
    defer func() {
        err := file.Close()
        if err != nil {
            log.Println(err)
        }
    }()
    if err != nil {
        log.Println(err)
        return false
    }
    _, err = file.Write(text)
    if err != nil {
        log.Println(err)
        return false
    }
    return true
}

func checkExistence(s string, dir *string) bool {
// checks if a given file exists, file may be removed or changed between checking and returning
    fileInfo, err := os.Stat(*dir + "/" + s)
    if err != nil && fileInfo.IsDir() {
        return false
    }
    return true
}
