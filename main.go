package main

import (
	"log"

	"chpid/cmdl"
)

func main() {
	err := cmdl.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err:%v", err)
	}
	// time.Sleep(time.Second * 10)
}
