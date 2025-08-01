package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"gtodo/todo"
	"log"
	"os"
	"strings"
)

func GetUserAdvice() (bool, error) {
	confirmMessage := "Need to create an empty \".todos.json\" file in your home directory to store your todo items, continue? (y/n): "
	fmt.Print(confirmMessage)

	r := bufio.NewReader(os.Stdin)
	str, _ := r.ReadString('\n')
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)

	switch str {
	case "y", "yes":
		return true, nil
	case "n", "no":
		return false, nil
	default:
		return false, errors.New("please input \"yes\" or \"no\"")
	}
}

func Reload(t *todo.Todos) {
	_, err := os.Stat(todo.Path)
	if err != nil {
		log.Fatal("please run init")
	} else {
		if err := t.Load(); err != nil {
			log.Fatal(err)
		}
	}
}
