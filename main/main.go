package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/zaiddkhan/todo-cli"
	"io"
	"os"
	"strings"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "Add a new tag")
	complete := flag.Int("complete", 0, "mark a todo as completed")
	del := flag.Int("del", 0, "Delete a todo")
	list := flag.Bool("list", false, "List all todos")
	flag.Parse()
	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		panic(err)
		os.Exit(1)
	}

	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			panic(err)
		}
		todos.Add(task)
		err2 := todos.Store(todoFile)
		if err2 != nil {
			panic(err2)
		}
	case *complete > 0:
		err := todos.Completed(*complete)
		if err != nil {
			panic(err)
		}
		err2 := todos.Store(todoFile)
		if err2 != nil {
			panic(err2)
		}
	case *del > 0:
		err := todos.Delete(*del)
		if err != nil {
			panic(err)
		}
		err2 := todos.Store(todoFile)
		if err2 != nil {
			panic(err2)
		}
	case *list:
		todos.Print()
	default:
		fmt.Fprintf(os.Stdout, "invalid command")
		os.Exit(0)

	}

}

func getInput(r io.ReadCloser, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	text := scanner.Text()
	if len(text) == 0 {
		return "", errors.New("empty input")
	}
	return text, nil
}
