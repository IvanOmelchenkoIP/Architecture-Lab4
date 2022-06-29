package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	
	engine "github.com/IvanOmelchenkoIP/Architecture-Lab4/engine"
)

func Parse(line string) engine.Command {
	var cmd engine.Command
	
	lineArgs := strings.Split(line, " ")
	command := lineArgs[0]
	switch command {
	case "print":
		arg := getArguments(lineArgs)
		cmd = engine.CommandPrint{Arg: arg}
	case "reverse":
		arg := getArguments(lineArgs)
		cmd = engine.CommandReverse{Arg: arg}
	default:
		errMessage := fmt.Sprintf("SYNTAX ERROR: %s", command)
		cmd = engine.CommandPrint{Arg: errMessage} 
	}
	return cmd
}

func getArguments(lineArgs []string) string {
	arg := strings.Join(lineArgs[1:], " ")
	return arg
}

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
