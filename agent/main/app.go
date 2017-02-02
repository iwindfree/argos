package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/iwindfree/argos/agent/manager"
)

func main() {
	var wg sync.WaitGroup
	command := os.Args[1]
	switch command {
	case "start":
		start(&wg)
	case "stop":
		stop()
	default:

	}

	manager.ServiceStart()
	wg.Wait()
	for {

	}
}

func start(workGroup *sync.WaitGroup) {
	displayLogo()
	workGroup.Add(1)
	displayLogo()
	manager.ServiceStart()

}

func stop() {

}

func displayLogo() {
	fmt.Println("MySQL Agent.")
}

func fileExist(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
