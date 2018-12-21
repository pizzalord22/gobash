package main

import (
    "gopkg.in/natefinch/lumberjack.v2"
    "log"
)

func setLogger(path, fileName string, size, backups, age int) {
    log.SetOutput(&lumberjack.Logger{
        Filename:   path + fileName,
        MaxSize:    50, // megabytes
        MaxBackups: 3,
        MaxAge:     28, //days
    })
}
