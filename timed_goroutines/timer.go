package main

import (
	"fmt"
	"time"
)

type timer struct {
	timer *time.Timer
	die   chan bool
	ping  chan string
}

func (t *timer) loop(expirancy int) {
	//defer die signal
	defer t.exit()

	for {
		select {
		case <-t.ping:
			fmt.Printf("Update timer\n")
			t.timer = getTimer(expirancy)
		case <-t.timer.C:
			fmt.Printf("Expired timer \n")
			return
		}
	}
}

func (t *timer) exit() {
	t.die <- true
}

func getTimer(expirancy int) *time.Timer {
	return time.NewTimer(time.Second * time.Duration(expirancy))
}
