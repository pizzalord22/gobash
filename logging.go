package main

import (
    "gopkg.in/natefinch/lumberjack.v2"
    "github.com/Sirupsen/logrus"
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

// todo switch to logrus: https://github.com/Sirupsen/logrus
//  the logrus library uses multiple logging errors to indicates severties this should be implemented together with the error struct

var logger = logrus.NewLogger()
