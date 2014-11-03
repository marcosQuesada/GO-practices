package main

import (
	"fmt"
)

type worker struct {
	id    int
	owner int
	read  chan string
	exit  chan int
}

func (w *worker) spawn(times int, done chan bool) {
	pointer := 0
	fmt.Printf("Spawnning id:%d \n", w.id)
	done <- true
	for {
		select {
		case msg := <-w.read:
			fmt.Printf("received: %s on worker %d \n", msg, w.id)
			pointer++
			if r.exists(w.owner) {
				fmt.Printf("Forwarding msg: %s to: %d iteration %d \n", msg, w.owner, pointer)
				fwdChan := r.getWorker(w.owner).read
				fwdChan <- msg
			}
			if pointer == times {
				fmt.Printf("Breaking, on %d\n", w.id)
				r.unregister(w.id)
				close(w.read)
				w.exit <- w.id
				return
			}
		}
	}
}
