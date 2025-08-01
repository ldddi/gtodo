package cmd

import (
	"fmt"
	"gtodo/todo"
	"log"
	"os"
)

// create data file
func Init() {
	ok, err := GetUserAdvice()
	if err != nil {
		log.Fatal(err)
	}

	if ok {
		_, err = os.Stat(todo.Path)
		if err != nil {
			if os.IsNotExist(err) {
				file, err := os.Create(todo.Path)
				if err != nil {
					log.Fatal(err)
				}
				defer file.Close()

				fmt.Printf("Succefully create .todos.json file in \"%v\"", todo.Path)
			} else {
				fmt.Print("Unkown error occured.")
			}
		} else {
			fmt.Printf("\".todos.json\" already exit! in %v", todo.Path)
		}
	}
}
