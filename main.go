package main

import (
	engine "github.com/IvanOmelchenkoIP/Architecture-Lab4/engine"
)

func main() {
	loop := new(engine.Loop)
	loop.Start()

	loop.AwaitFinish()
}
