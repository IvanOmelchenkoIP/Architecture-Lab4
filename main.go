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

type Queue struct {
	storage []Command
}

func (queue *Queue) push(cmd Command) {
	queue.storage = append(queue.storage, cmd)
}

func (queue *Queue) get() Command {
	if len(queue.storage) == 0 {
		return nil
	}
	cmd := queue.storage[0]
	queue.storage[0] = nil
	queue.storage = queue.storage[1:]
	return cmd
}

type Loop struct {
	commands *Queue
}

func (loop *Loop) Start() {
	for {
		cmd := loop.commands.get()
		if cmd != nil {
			cmd.Execute(loop)
		}
	}

}

func (loop *Loop) Post(cmd Command) {
}

func (loop *Loop) AwaitFinish() {

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
