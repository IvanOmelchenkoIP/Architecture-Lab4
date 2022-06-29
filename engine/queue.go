package engine

import "sync"

type eventQueue struct {
	mutex   sync.Mutex
	storage []Command

	isNotEmpty chan struct{}
	wait       bool
}

func (queue *eventQueue) push(cmd Command) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	queue.storage = append(queue.storage, cmd)

	if queue.wait {
		queue.wait = false
		queue.isNotEmpty <- struct{}{}
	}
}

func (queue *eventQueue) get() Command {
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

func (queue *eventQueue) empty() bool {
	return len(queue.storage) == 0
}
