package parser

import (
	"testing"
	gocheck "gopkg.in/check.v1"
	engine "github.com/IvanOmelchenkoIP/Architecture-Lab4/engine"
)

func Test(t *testing.T) { gocheck.TestingT(t) }

type MySuite struct{}

var _ = gocheck.Suite(&MySuite{})

func (s *MySuite) TestParserPrint(c *gocheck.C) {
	var input = "print hello"
	var expected = engine.CommandPrint{Arg: "hello"}
	cmd := Parse(input)
	c.Assert(cmd, gocheck.Equals, expected)
}


func (s *MySuite) TestParserPrintSpaceString(c *gocheck.C) {
	var input = "print hello world"
	var expected = engine.CommandPrint{Arg: "hello world"}
	cmd := Parse(input)
	c.Assert(cmd, gocheck.Equals, expected)
}

func (s *MySuite) TestParserPrintQuoteString(c *gocheck.C) {
	var input = "print \"string in quotes\""
	var expected = engine.CommandPrint{Arg: "\"string in quotes\""}
	cmd := Parse(input)
	c.Assert(cmd, gocheck.Equals, expected)
}

func (s *MySuite) TestParserPrintReverse(c *gocheck.C) {
	var input = "print reverse string"
	var expected = engine.CommandPrint{Arg: "reverse string"}
	cmd := Parse(input)
	c.Assert(cmd, gocheck.Equals, expected)
}

func (s *MySuite) TestParserReverse(c *gocheck.C) {
	var input = "reverse string"
	var expected = engine.CommandReverse{Arg: "string"}
	cmd := Parse(input)
	c.Assert(cmd, gocheck.Equals, expected)
}

func (s *MySuite) TestParserReverseSpaceString(c *gocheck.C) {
	var input = "reverse string to reverse"
	var expected = engine.CommandReverse{Arg: "string to reverse"}
	cmd := Parse(input)
	c.Assert(cmd, gocheck.Equals, expected)
}

func (s *MySuite) TestParserReverseQuoteString(c *gocheck.C) {
	var input = "reverse \"string in quotes to reverse\""
	var expected = engine.CommandReverse{Arg: "\"string in quotes to reverse\""}
	cmd := Parse(input)
	c.Assert(cmd, gocheck.Equals, expected)
}

func (s *MySuite) TestParserReversePrint(c *gocheck.C) {
	var input = "reverse print string"
	var expected = engine.CommandReverse{Arg: "print string"}
	cmd := Parse(input)
	c.Assert(cmd, gocheck.Equals, expected)
}

func (s *MySuite) TestParserSyntaxError(c *gocheck.C) {
	var input = "mul 3 4"
	var expected = engine.CommandPrint{Arg: "SYNTAX ERROR: mul"}
	cmd := Parse(input)
	c.Assert(cmd, gocheck.Equals, expected)
}

func (s *MySuite) TestParserSyntaxErrorPrintArg(c *gocheck.C) {
	var input = "add print string"
	var expected = engine.CommandPrint{Arg: "SYNTAX ERROR: add"}
	cmd := Parse(input)
	c.Assert(cmd, gocheck.Equals, expected)
}

func (s *MySuite) TestParserSyntaxErrorReverseArg(c *gocheck.C) {
	var input = "sub reverse string"
	var expected = engine.CommandPrint{Arg: "SYNTAX ERROR: sub"}
	cmd := Parse(input)
	c.Assert(cmd, gocheck.Equals, expected)
}

func (s *MySuite) TestParserSyntaxErrorPrintTypo(c *gocheck.C) {
	var input = "prinf string"
	var expected = engine.CommandPrint{Arg: "SYNTAX ERROR: prinf"}
	cmd := Parse(input)
	c.Assert(cmd, gocheck.Equals, expected)
}

func (s *MySuite) TestParserSyntaxErrorReverseTypo(c *gocheck.C) {
	var input = "reverce string"
	var expected = engine.CommandPrint{Arg: "SYNTAX ERROR: reverce"}
	cmd := Parse(input)
	c.Assert(cmd, gocheck.Equals, expected)
}