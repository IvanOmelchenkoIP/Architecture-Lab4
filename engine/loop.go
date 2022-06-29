package engine

type Loop struct {
	commands *eventQueue

	stop       bool
	stopSignal chan struct{}
}

func (loop *Loop) Start() {
	loop.commands = &eventQueue{
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
	// Commands after the request to stop will be ignored
	if loop.stop {
		return
	}
	loop.commands.push(cmd)
}

func (loop *Loop) AwaitFinish() {
	loop.Post(commandStop(func(handler Handler) {
		loop.stop = true
	}))
	<-loop.stopSignal
}

type commandStop func(handler Handler)

func (stop commandStop) Execute(handler Handler) {
	stop(handler)
}
