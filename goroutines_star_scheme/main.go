package main

import (
	"flag"
	"fmt"
)

type register struct {
	workers map[int]worker
}

func main() {
	processes := flag.Int("goroutines", 1, "number of goroutines")
	times := flag.Int("times", 1, "ring iteration")
	flag.Parse()

	fmt.Printf("foo processes %d times %d \n", *processes, *times)
	var r = register{
		workers: make(map[int]worker),
	}

	for i := 1; i <= *processes; i++ {
		var w = worker{
			id:   i,
			read: make(chan string, 0),
			end:  make(chan bool, 0),
		}
		r.workers[i] = w
		go w.spawn()
	}

	for {
		var command string
		fmt.Scanf("%s \n", &command)
		switch command {
		case "kill":
			var process int
			fmt.Scanf("%d \n", &process)
			if r.exists(process) {
				r.workers[process].end <- true
				r.unregister(process)
			} else {
				fmt.Printf("Goroutine not found \n")
			}
		case "exit":
			fmt.Println("exiting")
			return
		default:
			var process int
			fmt.Scanf("%d \n", &process)
			fmt.Printf("command:%s on Process: %d\n", command, process)
			if r.exists(process) {
				r.workers[process].read <- command
			} else {
				fmt.Printf("Goroutine not found \n")
			}
		}
	}
}

func (r *register) exists(id int) bool {
	_, ok := r.workers[id]
	return ok
}

func (r *register) unregister(id int) {
	delete(r.workers, id)
}
