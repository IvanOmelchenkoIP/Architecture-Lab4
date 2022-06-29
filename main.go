package main

import (
	"fmt"
	"sync"
)

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(cmd Command)
}

type Queue struct {
	mutex   sync.Mutex
	storage []Command

	isNotEmpty chan struct{}
	wait       bool
}

func (queue *Queue) push(cmd Command) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	queue.storage = append(queue.storage, cmd)

	if queue.wait {
		queue.wait = false
		queue.isNotEmpty <- struct{}{}
	}
}

func (queue *Queue) get() Command {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	if queue.empty() {
		queue.mutex.Unlock()
		queue.wait = true
		<-queue.isNotEmpty
		queue.mutex.Lock()
	}

	cmd := queue.storage[0]
	queue.storage[0] = nil
	queue.storage = queue.storage[1:]
	return cmd
}

func (queue *Queue) empty() bool {
	return len(queue.storage) == 0
}

type Loop struct {
	commands *Queue

	stop       bool
	stopSignal chan struct{}
}

func (loop *Loop) Start() {
	loop.commands = &Queue{
		storage:    make([]Command, 0),
		isNotEmpty: make(chan struct{}),
	}
	loop.stopSignal = make(chan struct{})

	go func() {
		for !loop.stop || !loop.commands.empty() {
			cmd := loop.commands.get()
			cmd.Execute(loop)
		}
		loop.stopSignal <- struct{}{}
	}()
}

func (loop *Loop) Post(cmd Command) {
	loop.commands.push(cmd)
}

type CommandStop struct{}

func (stop CommandStop) Execute(handler Handler) {
	handler.(*Loop).stop = true
}

func (loop *Loop) AwaitFinish() {
	loop.Post(CommandStop{})
	<-loop.stopSignal
}

type CommandPrint struct {
	arg string
}

func (print CommandPrint) Execute(handler Handler) {
	fmt.Println(print.arg)
}

type CommandReverse struct {
	arg string
}

func (reverse CommandReverse) Execute(handler Handler) {
	strBytes := []rune(reverse.arg)
	reversedBytes := make([]rune, 0)
	for i := len(strBytes) - 1; i >= 0; i-- {
		reversedBytes = append(reversedBytes, strBytes[i])
	}
	reversed := string(reversedBytes)
	handler.Post(&CommandPrint{arg: reversed})
}

func main() {
	loop := new(Loop)
	loop.Start()

	loop.AwaitFinish()
}
