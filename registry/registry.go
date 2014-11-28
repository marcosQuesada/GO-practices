package registry

import (
	"fmt"
	"strings"
	"time"
)

type registry struct {
	index      map[string]string
	inputChan  chan string
	outputChan chan string
	stop       bool
}

var inputChannel = make(chan string, 1)
var outputChannel = make(chan string, 1)

func Start() {

	registry := &registry{
		index:      make(map[string]string),
		inputChan:  inputChannel,
		outputChan: outputChannel,
		stop:       false,
	}

	go registry.run()
}

func Set(key string, value string) {
	inputChannel <- fmt.Sprintf("set %s %s", key, value)
}

func Get(key string) string {
	timeout := time.NewTimer(time.Second * 1)
	inputChannel <- fmt.Sprintf("get %s", key)
	select {
	case response := <-outputChannel:
		return response
	case <-timeout.C:
		fmt.Printf("Get %s Timeout \n", key)
		return ""
	}
}

func (r *registry) run() {
	for !r.stop {
		select {
		case msg := <-r.inputChan:
			msgParts := strings.Split(msg, " ")
			switch msgParts[0] {
			case "stop":
				r.stop = true
			case "set":
				if len(msgParts) != 3 {
					fmt.Printf("Set Bad Format: %s \n", len(msgParts))
				}
				r.index[msgParts[1]] = msgParts[2]

			case "get":
				if len(msgParts) != 2 {
					fmt.Printf("Bad Format: %s \n", msg)
				}
				value, exists := r.index[msgParts[1]]

				if !exists {
					fmt.Printf("key not found:%s\n", msgParts[1])
				}

				r.outputChan <- value
			}
		}
	}
}
