package main

import (
    "gopkg.in/natefinch/lumberjack.v2"
    "log"
)

func setLogger(path, fileName string, size, backups, age int) {
    log.SetOutput(&lumberjack.Logger{
        Filename:   path + fileName,
        MaxSize:    size, // megabytes
        MaxBackups: backups,
        MaxAge:     age, //days
    })
}
