package main

import (
	"flag"
	"gtodo/cmd"
	"gtodo/todo"
	"os"
)

func main() {
	todos := &todo.Todos{}

	flag.Parse()
	f := flag.Arg(0)

	switch f {
	case "":

	case "help":
		cmd.Help()
	case "init":
		cmd.Init()
	case "add":
		cmd.Reload(todos)
		cmd.AddTask(todos, os.Args[:2])
	case "list":
		cmd.AddTask(todos, os.Args[:2])
		cmd.Reload(todos)
	}
}
