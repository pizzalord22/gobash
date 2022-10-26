package main

import (
    "gopkg.in/natefinch/lumberjack.v2"
    "log"
)

// setup the logger
func setLogger(path, fileName string, backups, age int) {
    log.SetOutput(&lumberjack.Logger{
        Filename:   path + fileName,
        MaxSize:    5, // megabytes
        MaxBackups: backups,
        MaxAge:     age, //days
    })
}
