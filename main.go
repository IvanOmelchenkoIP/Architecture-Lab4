package main

import ( "fmt" )

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(cmd Command)
}

func main() {

}