package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	writePid()
}

func writePid() error {
	path, err := os.Getwd()
	if err == nil {
		pidfile := path + "/" + strconv.Itoa(os.Getpid()) + ".scouter"
		fmt.Printf("pidfile: %s\n", pidfile)
		fmt.Println("start sleep")
		time.Sleep(1 * time.Second)
		fmt.Println("end sleep")
		os.Create(pidfile)
	} else {
		return err
	}
	return nil
}
