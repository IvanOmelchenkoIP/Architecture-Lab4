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
	loop.commands.push(cmd)
}

func (loop *Loop) AwaitFinish() {
	loop.Post(commandStop{})
	<-loop.stopSignal
}

type commandStop struct{}

func (stop commandStop) Execute(handler Handler) {
	handler.(*Loop).stop = true
}