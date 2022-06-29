package main

import (
	"fmt"
)

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(cmd Command)
}

type CommandPrint struct {
	str string
}

func (print *CommandPrint) Execute(handler Handler) {
	fmt.Println(print.str)
}

type CommandReverse struct {
	str string
}

func (reverse *CommandReverse) Execute(handler Handler) {
	strBytes := []rune(reverse.str)
	reversedBytes := make([]rune, 0)
	for i := len(strBytes) - 1; i >= 0; i-- {
		reversedBytes = append(reversedBytes, strBytes[i])
	}
	reversedStr := string(reversedBytes)
	handler.Post(&CommandPrint{str: reversedStr})
}

func main() {

}
