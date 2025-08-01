package todo

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"time"
)

// todo struct
type item struct {
	ID          int
	Task        string
	Category    string
	Done        bool
	CreateAt    time.Time
	CompletedAt *time.Time // 节约内存，未完成时为nil
}

// todo list
type Todos []item

var Path string

var nextID int

// init store file path
func init() {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	Path = filepath.Join(dir, ".todos.json")
}

func (t *Todos) getIdxByID(id int) int {
	idx := -1
	for i, item := range *t {
		if item.ID == id {
			idx = i
			break
		}
	}

	return idx
}

func (t *Todos) getNotDoneCnt() int {
	cnt := 0
	for _, item := range *t {
		if !item.Done {
			cnt++
		}
	}

	return cnt
}

func (t *Todos) Add(task, cat string) {
	item := item{
		ID:          nextID,
		Task:        task,
		Category:    cat,
		Done:        false,
		CreateAt:    time.Now(),
		CompletedAt: nil,
	}

	*t = append(*t, item)
}

func (t *Todos) Update(id int, task string, cat string, done int) error {
	idx := t.getIdxByID(id)

	if idx < 0 {
		return errors.New("Invalid ID")
	}

	if task != "" {
		(*t)[idx].Task = task
	}

	if cat != "" {
		(*t)[idx].Category = cat
	}

	switch done {
	case 0:
		(*t)[idx].Done = false
		(*t)[idx].CompletedAt = nil
	case 1:
		now := time.Now()
		(*t)[idx].Done = true
		(*t)[idx].CompletedAt = &now
	}
	return nil
}

func (t *Todos) Delete(id int) error {
	todos := *t

	idx := t.getIdxByID(id)

	if idx < 0 {
		return errors.New("invalid ID")
	}

	*t = append(todos[:idx], todos[idx+1:]...)
	return nil
}

func (t *Todos) Load() error {
	fileBytes, err := os.ReadFile(Path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(fileBytes, t); err != nil {
		return err
	}

	nextID = (*t)[0].ID
	for _, item := range *t {
		if nextID > item.ID {
			nextID = item.ID + 1
		}
	}
	return nil
}

func (t *Todos) Store() error {

	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(Path, data, 0644)
}

func (t *Todos) Print() {

}
