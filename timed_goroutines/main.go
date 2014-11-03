package main

import (
	"flag"
	"fmt"
)

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
