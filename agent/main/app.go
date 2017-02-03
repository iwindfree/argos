package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/iwindfree/argos/agent/manager"
)

var pidfile string
var f *os.File

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
	writePid()

	for fileExist(pidfile) {
		time.Sleep(1 * time.Second)
	}

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

func writePid() error {
	path, err := os.Getwd()
	if err == nil {
		pidfile = path + "/" + strconv.Itoa(os.Getpid()) + ".scouter"
		f, err = os.Create(pidfile)
		if err != nil {
			return err
		}
		defer f.Close()
	} else {
		return nil
	}
	return nil
}
