package cmd

import (
	"flag"
	"gtodo/todo"
	"log"
)

func UpdateTask(t *todo.Todos, args []string) {
	fSet := flag.NewFlagSet("update", flag.ExitOnError)
	id := fSet.Int("id", -1, "ID of the task to update")
	task := fSet.String("task", "", "task description")
	cat := fSet.String("cat", "", "category of the task")
	done := fSet.Int("done", -1, "mark task as done or not done (1 for done, 0 for not done)")

	fSet.Parse(args)

	if err := t.Update(*id, *task, *cat, *done); err != nil {
		log.Fatal(err)
	}

	if err := t.Store(); err != nil {
		log.Fatal(err)
	}
}
