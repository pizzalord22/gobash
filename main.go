package main

import (
    "bufio"
    "fmt"
    "github.com/spf13/viper"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"
)

const bufferSize = 100

func init() {
    viper.AddConfigPath("./")
    viper.SetConfigName("defaults")
    err := viper.ReadInConfig()
    if err != nil{
        fmt.Println(err)
        return
    }
    logPath := viper.GetString("logPath")
    logName := viper.GetString("logName")
    logSize := viper.GetInt("logSize")
    logMaxbackups := viper.GetInt("maxBackUps")
    MaxAge := viper.GetInt("MaxAge")
    fmt.Println("LogPath:", logPath)
    fmt.Println("LogName:", logName)
    err = os.MkdirAll(logPath, os.ModePerm)
    CheckError(err, logPath)
    f, err := os.Create(fmt.Sprintf("%s/%s",logPath,logName))
    defer f.Close()
	setLogger(logPath,logName,logSize,logMaxbackups,MaxAge)
    if err != nil {
        return
    }
    go webServer()
}

// main loop here we initialize our reader
func main() {

    reader := bufio.NewReader(os.Stdin)
    ex, err := os.Executable()
    if err != nil {
        fmt.Println(err)
        fmt.Println("exiting the simulation in 5 seconds")
        time.Sleep(5 * time.Second)
        return
    }
    CheckError(err)
    dir := strings.Replace(filepath.Dir(ex), "\\", "/", -1)
    fmt.Printf("%v>", dir)
    osLoop(reader, &dir)
}

// this simulates an os loop
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
        case "cat":
            if len(r) > 2 {
                cat(r[1], r[2], r[3])
            } else {
                showHelp("cat")
            }
        case "webdog":
            execute([]string{"start", "chrome.exe", "mydogchase.com"})
        case "webstat":
            if len(r) > 1 {
                startwebserver(r[1])
            }else{
                startwebserver("8080")
            }
        default:
            log.Printf("Command: %s not found", r[0])
            fmt.Printf("gobash> Type \"help\" to see a list of available commands\n")
        }
        CheckError(err)
        fmt.Printf("%v>", *dir)
    }
}
