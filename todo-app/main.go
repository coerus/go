package todo

import (
	"fmt"
	"encoding/json"
	"time"
	"os"
	"io/ioutil"
	"errors"
)

type item struct {
	Task string
	Done bool
	CreatedAt time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(task string){
	t := item {
		Task:  task,
		Done: false,
		CreatedAt: time.Now(),
		CompletedAt: time.Time{},
	}
 *l = append(*l, t)
}

func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d não existe", i)

	}
ls[i-1].Done = true
ls[i-1].CompletedAt = time.Now()

return nil
}

func (l *List) Save(fileName string) error {
	json, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, json, 0644)
}

func (l *List) Get(fileName string) error {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		
        if errors.Is(err, os.ErrNotExist) {
            return nil
        }
        return err

	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}
