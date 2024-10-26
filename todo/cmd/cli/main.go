package main

import (
	"fmt"
	"os"
	"strings"

	"github.com.Shashwat1977.todo/internal/model"
)

func main() {
	l := &model.List{}

	const todoFileName = ".todo.json"

	if err := l.Get(todoFileName); err != nil { // If the file does not exist it does not throw an error.
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
