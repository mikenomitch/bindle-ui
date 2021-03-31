package main

import (
	"log"
	"os/exec"

	"github.com/mikenomitch/bindle-packager/service"
)

func main() {
	cmd := exec.Command("bindle", "init")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	service.Run()
}
