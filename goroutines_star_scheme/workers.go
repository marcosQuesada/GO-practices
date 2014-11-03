package main

import (
	"fmt"
)

type worker struct {
	id   int
	read chan string
	end  chan bool
}

func (w *worker) spawn() {
	fmt.Printf("Spawnning id:%d \n", w.id)
	for {
		select {
		case msg := <-w.read:
			fmt.Printf("received: %s on worker %d \n", msg, w.id)
		case <-w.end:
			fmt.Printf("Breaking, on %d\n", w.id)
			close(w.end)
			return
		}
	}
}
