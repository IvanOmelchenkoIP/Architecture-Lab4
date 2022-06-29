package parser

import (
	"fmt"
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
