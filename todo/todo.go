package todo

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
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

func (t *Todos) DeleteByID(id int) error {

	idx := t.getIdxByID(id)
	if idx < 0 {
		return errors.New("invalid ID")
	}

	*t = append((*t)[:idx], (*t)[idx+1:]...)
	fmt.Println("Successful")
	return nil
}

func (t *Todos) DeleteByDone(done int) error {

	if done != 0 && done != 1 {
		return errors.New("invalid ID")
	}

	if done == 0 {
		fmt.Print("This will delete the unfinished todos, are you sure? (yes/no)")
		r := bufio.NewReader(os.Stdin)
		ok, err := r.ReadString('\n')
		if err != nil {
			return err
		}
		ok = strings.TrimSpace(ok)
		ok = strings.ToLower(ok)

		if ok == "no" {
			return nil
		}
	}

	tmp := (*t)[:0]
	for _, item := range *t {
		if done == 1 && !item.Done {
			tmp = append(tmp, item)
		} else if done == 0 && item.Done {
			tmp = append(tmp, item)
		}
	}

	*t = tmp

	fmt.Println("Successful")
	return nil
}

func (t *Todos) Load() error {
	fileBytes, err := os.ReadFile(Path)
	if err != nil {
		return err
	}

	if len(fileBytes) == 0 {
		return err
	}

	if err := json.Unmarshal(fileBytes, t); err != nil {
		return err
	}

	nextID = -1
	for _, item := range *t {
		if item.ID > nextID {
			nextID = item.ID
		}
	}
	nextID++
	return nil
}

func (t *Todos) Store() error {

	data, err := json.MarshalIndent(t, "", "")
	if err != nil {
		return err
	}

	return os.WriteFile(Path, data, 0644)
}

func (t *Todos) Print() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Text: "ID", Align: simpletable.AlignCenter},
			{Text: "Task", Align: simpletable.AlignCenter},
			{Text: "Category", Align: simpletable.AlignCenter},
			{Text: "Done", Align: simpletable.AlignCenter},
			{Text: "CreateAt", Align: simpletable.AlignCenter},
			{Text: "CompletedAt", Align: simpletable.AlignCenter},
		},
	}

	for _, row := range *t {
		done := "❌"
		completeAt := ""

		if row.Done {
			done = "\u2705"
		}

		if row.CompletedAt != nil {
			completeAt = row.CompletedAt.Format("2006-01-02")
		}

		table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", row.ID), Align: simpletable.AlignCenter},
			{Text: row.Task, Align: simpletable.AlignCenter},
			{Text: row.Category, Align: simpletable.AlignCenter},
			{Text: done, Align: simpletable.AlignCenter},
			{Text: row.CreateAt.Format("2006-01-02"), Align: simpletable.AlignCenter},
			{Text: completeAt, Align: simpletable.AlignCenter},
		})
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Text: fmt.Sprintf("not done count: %v", t.getNotDoneCnt()), Span: 5},
			{},
		},
	}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}
