package main

import (
    "github.com/DGHeroin/event.go"
    "log"
)

type HelloListener int

func (p *HelloListener) OnEvent(args interface{}) {
    log.Println("===>", args, *p)
}

func main()  {
    storage := event.NewMemoryStorage()
    dispatcher := event.NewDispatcher()
    mgr := event.NewManager(storage, dispatcher, nil)


    var lis HelloListener = 1001
    mgr.Add("hello", &lis)

    mgr.Dispatch("hello", "world")

}
