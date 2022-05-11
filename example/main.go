package main

import (
    "github.com/DGHeroin/event.go"
    "log"
)

type HelloListener int

func (p *HelloListener) OnEvent(args event.F) {
    log.Println("===>", args, *p)
}

func T1()  {
    storage := event.NewMemoryStorage()
    dispatcher := event.NewDispatcher()
    mgr := event.NewManager(storage, dispatcher, nil)

    var lis HelloListener = 1001
    mgr.Add("hello", &lis)

    mgr.Dispatch("hello", event.F{
        "args": "world",
    })
}
func T2()  {
    eventName := "event-name"
    ln1 := event.On(eventName, func(f event.F) {
        log.Println("1. process:", f)
    })
    ln2 := event.On(eventName, func(f event.F) {
        log.Println("2. process:", f)
    })
    event.Emit(eventName, event.F{
        "key": "hello!",
    })

    event.Off(eventName, ln1)

    event.Emit(eventName, event.F{
        "key": "sayonara!",
    })
    event.Off(eventName, ln1)
    event.Off(eventName, ln2)
}
func main()  {
    T1()
    T2()
}
