package cmd

import (
	"flag"
	"gtodo/todo"
)

func ListTask(t *todo.Todos, args []string) {
	fset := flag.NewFlagSet("list", flag.ExitOnError)
	done := fset.Int("done", -1, "show done(1) or not done(0)")
	cat := fset.String("cat", "", "filter by category")

	fset.Parse(args)

	tmp := (*t)[:0]
	for _, item := range *t {
		if *done == 1 && item.Done {
			if *cat != "" && item.Category == *cat {
				tmp = append(tmp, item)
			} else if *cat == "" {
				tmp = append(tmp, item)
			}
		} else if *done == 0 && !item.Done {
			if *cat != "" && item.Category == *cat {
				tmp = append(tmp, item)
			} else if *cat == "" {
				tmp = append(tmp, item)
			}
		}
	}

	if len(tmp) != 0 {
		t = &tmp
	}

	t.Print()
}
