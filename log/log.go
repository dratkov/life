package log

import (
    "os"
    "log"
)

type Log struct {
    filename string
    file_handler *os.File
}

func (l *Log) Write( text string )  {
    if l.file_handler == nil {
        f, err := os.OpenFile(l.filename, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666)
        if err != nil {

        }
        l.file_handler = f
    }

    log.SetOutput(l.file_handler)
    log.Println(text)
}

func (l *Log) Destruct() {
    l.file_handler.Close()
    l.file_handler = nil
}

func (l *Log) SetFileName( filename string )  {
    l.filename = filename
}

var log_singletone *Log = nil

func New() *Log  {
    if log_singletone != nil {
        return log_singletone
    }
    log_singletone = &Log{}
    log_singletone.SetFileName(".logs/life.log")

    return log_singletone
}
