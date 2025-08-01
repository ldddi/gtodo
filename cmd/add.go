package cmd

import (
	"flag"
	"gtodo/todo"
	"log"
)

func AddTask(t *todo.Todos, args []string) {
	fset := flag.NewFlagSet("add", flag.ExitOnError)
	task := fset.String("task", "", "The content of new todo item")
	cat := fset.String("cat", "", "The category of new todo item")

	fset.Parse(args)

	if *task == "" {
		log.Fatal("Error: -task flag is required")
	}

	if *cat == "" {
		*cat = "default"
	}

	t.Add(*task, *cat)
}
