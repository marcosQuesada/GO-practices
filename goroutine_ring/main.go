package main

import (
	"flag"
	"fmt"
)

func main() {
	processes := flag.Int("goroutines", 1, "number of goroutines")
	times := flag.Int("times", 1, "ring iteration")
	flag.Parse()

	fmt.Printf("Processes %d times %d \n", *processes, *times)
	//finish channel
	finish := make(chan int, 1)
	//create ring
	for i := 1; i < *processes; i++ {
		var w = worker{
			id:    i,
			owner: i + 1,
			read:  make(chan string, 0),
			exit:  finish,
		}
		done := make(chan bool, 1)
		go w.spawn(*times, done)
		r.register(i, w)
		<-done
	}

	//last element closes the ring
	var w = worker{
		id:    *processes,
		read:  make(chan string, 0),
		owner: 1,
		exit:  finish,
	}
	done := make(chan bool, 1)
	go w.spawn(*times, done)
	r.register(*processes, w)
	<-done

	//start message propagation
	readChan := r.getWorker(1).read
	readChan <- "init"

	//wait until last goroutine has exited
	fmt.Printf("Wait finish Signal")
	for {
		select {
		case msg := <-finish:
			fmt.Printf("finish %d \n", msg)
			if msg == *processes {
				fmt.Printf("Done \n")
				return
			}
		}
	}
}
