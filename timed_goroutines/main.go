package main

import (
	"flag"
	"fmt"
	"time"
)

type timer struct {
	timer *time.Timer
	die   chan bool
	ping  chan string
}

func main() {
	expires := flag.Int("expires", 1, "expirancy")
	flag.Parse()
	fmt.Printf("Init \n")
	var t = timer{
		timer: getTimer(*expires),
		die:   make(chan bool, 1),
		ping:  make(chan string, 0),
	}
	go t.loop(*expires)

	for {
		var msg string
		fmt.Scanf("%s \n", &msg)
		t.ping <- msg

		select {
		case <-t.die:
			fmt.Print("time to die \n")
			return
		}
	}
}

func (t *timer) loop(expirancy int) {
	for {
		select {
		case <-t.ping:
			fmt.Printf("Update timer\n")
			t.timer = getTimer(expirancy)
		case <-t.timer.C:
			fmt.Printf("Expired timer \n")
			t.die <- true
			return
		}
	}
}

func getTimer(expirancy int) *time.Timer {
	return time.NewTimer(time.Second * time.Duration(expirancy))
}
