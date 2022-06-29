package engine

import "fmt"

type CommandPrint struct {
	Arg string
}

func (print CommandPrint) Execute(handler Handler) {
	fmt.Println(print.Arg)
}

type CommandReverse struct {
	Arg string
}

func (reverse CommandReverse) Execute(handler Handler) {
	strBytes := []rune(reverse.Arg)
	reversedBytes := make([]rune, 0)
	for i := len(strBytes) - 1; i >= 0; i-- {
		reversedBytes = append(reversedBytes, strBytes[i])
	}
	reversed := string(reversedBytes)
	handler.Post(&CommandPrint{Arg: reversed})
}