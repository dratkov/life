package main

import (
    "./input"
    "./output"
    //"./strategy"
    "time"
    "syscall"
    "./log"
    //"fmt"
)

func  main() {
    l := log.New()
    defer l.Destruct()
    in := input.NewInput()
    in.InitFromSource()
    build := in.GetBuild()
    area := build.GetArea()
    //strategy := strategy.New( area )

    out := output.NewOutput( in )

    for i := 0; i < 3000; i++ {
        out.DoOutput()
        //fmt.Println( area, "++" )
        if area.GetLiveCellCount() == 0 {
            syscall.Exit(0)
        }
        time.Sleep(200 * time.Millisecond)
        area.NextTime()
    }

    syscall.Exit(0)
    //out.DoOutput()
/*
    time.Sleep(1 * time.Second)
    area.NextTime()
    out.DoOutput()

    time.Sleep(1 * time.Second)
    area.NextTime()
    out.DoOutput()

    time.Sleep(1 * time.Second)
    area.NextTime()
    out.DoOutput()
*/
}