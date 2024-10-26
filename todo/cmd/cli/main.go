package main

import (
	"flag"
	"fmt"
	"os"

	"github.com.Shashwat1977.todo/internal/model"
)

func main() {
	l := &model.List{}

	const todoFileName = ".todo.json"

	task := flag.String("task", "", "Task to include in todo list")
	list := flag.Bool("list", false, "List All Tasks")
	complete := flag.Int("complete", 0, "Mark a task as completed")

	flag.Parse()

	if err := l.Get(todoFileName); err != nil { // If the file does not exist it does not throw an error.
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		fmt.Print(l)
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		l.Add(*task)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Println("Invalid flag")
	}
}
