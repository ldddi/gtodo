package cmd

import (
	"flag"
	"gtodo/todo"
	"log"
)

func DeleteTask(t *todo.Todos, args []string) {
	fSet := flag.NewFlagSet("delete", flag.ExitOnError)
	id := fSet.Int("id", -1, "ID of the task to delete")
	done := fSet.Int("done", -1, "done status of the task to delete")

	fSet.Parse(args)

	if *id != -1 && *done == -1 {
		if err := t.DeleteByID(*id); err != nil {
			log.Fatal(err)
		}
	} else if *id == -1 && *done != -1 {
		if err := t.DeleteByDone(*done); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("err: please input id or done to delete")
	}

	if err := t.Store(); err != nil {
		log.Fatal(err)
	}
}
