package main

import (
    "fmt"
    "github.com/fatih/color"
    "os/exec"
    "runtime"
)

const windows = "windows"

// execute commands like it is a real os
func execute(command []string) {
    var s string
    osName := runtime.GOOS
    for _, v := range command {
        s += v + " "
    }
    if osName == windows {
        execWindows(s)
    } else {
        execSane(s)
    }
}

// executes commands on a windows machine
func execWindows(s string) {
    result, err := exec.Command("cmd", "/C", s).CombinedOutput()
    if err != nil {
        color.Red(fmt.Sprintf("gobash> The following error appeared when executing a command; %v\n", err))
        return
    }
    color.Green("gobash> executed: %v", s)
    if len(result) > 0 {
        fmt.Println(string(result))
    }
}

// executes commands the sane way
func execSane(s string) {
    result, err := exec.Command("/bin/sh", "-c", s).CombinedOutput()
    if err != nil {
        color.Red(fmt.Sprintf("gobash> The following error appeared when executing a command; %v\n", err))
        return
    }
    color.Green("gobash> executed: %v", s)
    if len(result) > 0 {
        fmt.Println(string(result))
    }
}
