package engine

import (
	"testing"
	gocheck "gopkg.in/check.v1"
)

func Test(t *testing.T) { gocheck.TestingT(t) }

type MySuite struct{}

var _ = gocheck.Suite(&MySuite{})

func (s *MySuite) TestLoopQueueLenPrint(c *gocheck.C) {
	var expected = 2
	loop := new(Loop)
	loop.Start()

	loop.Post(&CommandPrint{Arg: "str1"})
	loop.Post(&CommandPrint{Arg: "str2"})

	queueLen := len(loop.commands.storage)

	loop.AwaitFinish()
	c.Assert(queueLen, gocheck.Equals, expected)
}

func (s *MySuite) TestLoopQueueLenReverse(c *gocheck.C) {
	var expected = 3
	loop := new(Loop)
	loop.Start()

	loop.Post(&CommandReverse{Arg: "str1"})
	loop.Post(&CommandReverse{Arg: "str2"})
	loop.Post(&CommandReverse{Arg: "str3"})

	queueLen := len(loop.commands.storage)

	loop.AwaitFinish()
	c.Assert(queueLen, gocheck.Equals, expected)
}

func (s *MySuite) TestLoopQueueLenMixed(c *gocheck.C) {
	var expected = 4
	loop := new(Loop)
	loop.Start()

	loop.Post(&CommandReverse{Arg: "str1"})
	loop.Post(&CommandPrint{Arg: "str2"})
	loop.Post(&CommandReverse{Arg: "str3"})
	loop.Post(&CommandPrint{Arg: "str4"})

	queueLen := len(loop.commands.storage)

	loop.AwaitFinish()
	c.Assert(queueLen, gocheck.Equals, expected)
}

func (s *MySuite) TestLoopStopFlag(c *gocheck.C) {
	var expected = true
	loop := new(Loop)
	loop.Start()

	loop.Post(&CommandReverse{Arg: "str1"})
	loop.Post(&CommandPrint{Arg: "str2"})

	loop.AwaitFinish()

	stop := loop.stop
	c.Assert(stop, gocheck.Equals, expected)
}

func (s *MySuite) TestLoopStopAdd(c *gocheck.C) {
	var expected = 0
	loop := new(Loop)
	loop.Start()

	loop.Post(&CommandReverse{Arg: "str1"})
	loop.Post(&CommandPrint{Arg: "str2"})

	loop.AwaitFinish()

	loop.Post(&CommandReverse{Arg: "str1"})
	loop.Post(&CommandPrint{Arg: "str2"})

	queueLen := len(loop.commands.storage)
	c.Assert(queueLen, gocheck.Equals, expected)
}
