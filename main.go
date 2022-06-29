package main

import (
	"fmt"
	"os"
	"bufio"

	engine "github.com/IvanOmelchenkoIP/Architecture-Lab4/engine"
	. "github.com/IvanOmelchenkoIP/Architecture-Lab4/parser"
)

func main() {
	loop := new(engine.Loop)
	loop.Start()

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		errMessage := fmt.Sprintf("FILE READING ERROR: %s", filename)
		loop.Post(&engine.CommandPrint{Arg: errMessage})
	} else {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			cmd := Parse(line)
			loop.Post(cmd)
		}
	}

	loop.AwaitFinish()
}
