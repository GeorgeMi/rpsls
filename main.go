package main

import (
	"log"
	"os"

	"github.com/GeorgeMi/rpsls-api/commands"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("rpsls-api", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"run": commands.NewServerCommand,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Panic(err)
	}

	os.Exit(exitStatus)
}
