package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// indentation holds the current indentation level for String output.
// It is currently used to track the indentation of subelements, example:
// CheckCommand holds one or more CheckCommandArguments that need to be indented
// depending on their CheckCommand.
var indentation int

type Identifier string

// https://icinga.com/docs/icinga-2/latest/doc/17-language-reference/#duration-literals
type Duration struct {
	time.Duration
}

// https://icinga.com/docs/icinga-2/latest/doc/18-library-reference/#number-type
// Only Integers for now
type Integer int
type Float float64

// https://icinga.com/docs/icinga-2/latest/doc/18-library-reference/#string-type
type String string

// https://icinga.com/docs/icinga-2/latest/doc/18-library-reference/#dictionary-type
type Dictionary map[Identifier]Object

// https://icinga.com/docs/icinga-2/latest/doc/18-library-reference/#object-type
type Array []Object

// https://icinga.com/docs/icinga-2/latest/doc/18-library-reference/#boolean-type
type Boolean bool

// https://icinga.com/docs/icinga-2/latest/doc/18-library-reference/#object-type
type Object interface {
	String() string
}

// https://icinga.com/docs/icinga-2/latest/doc/18-library-reference/#type-type
// TODO Not yet implemented

// https://icinga.com/docs/icinga-2/latest/doc/18-library-reference/#datetime-type
// TODO  Not yet implemented

const (
	True  = Boolean(true)
	False = Boolean(false)
)

type Operator uint

const (
	PLUS Operator = iota
)

type InfixExpression struct {
	Left          Object
	InfixOperator Operator
	Right         Object
}

func (i Integer) String() string {
	return fmt.Sprintf("%d", i)
}

func (f Float) String() string {
	return fmt.Sprintf("%g", f)
}

// Wrapper to stringify the Icinga2 String Object
// Handles escaping and multiline strings when necessary
func (is String) String() string {
	if !strings.Contains(string(is), "\n") {
		// TODO Escape special characters properly
		return strconv.Quote(string(is))
	}

	if strings.Contains(string(is), "{{{") || strings.Contains(string(is), "}}}") {
		panic("Cannot properly escape string")
	}

	return `{{{` + string(is) + `}}}`
}

// RawString returns the String Object as string without escaping
func (is String) RawString() string {
	return string(is)
}

// String returns the Identifier Object as string
func (i Identifier) String() string {
	return string(i)
}

// String returns the Array Object as string.
// Uses [] as array markers and , as element delimiter
func (ia *Array) String() string {
	var b strings.Builder

	b.WriteString("[")

	if len(*ia) > 1 {
		for i := 0; i < len(*ia)-1; i++ {
			b.WriteString((*ia)[i].String() + ", ")
		}
		b.WriteString((*ia)[len(*ia)-1].String())
	} else if len(*ia) == 1 {
		b.WriteString((*ia)[0].String())
	}

	b.WriteString("]")

	return b.String()
}

// indentString is a helper function that returns tab indentation
func indentString() string {
	return strings.Repeat("\t", indentation)
}

// String returns the Operator Object as string
func (op *Operator) String() string {
	if *op == PLUS {
		return "+"
	}

	return ""
}

// String returns the InfixExpression Object as string
func (ie InfixExpression) String() string {
	var result strings.Builder

	result.WriteString(ie.Left.String() + " ")
	result.WriteString(ie.InfixOperator.String() + " ")
	result.WriteString((ie.Right).String())

	return result.String()
}

// String returns the Boolean Object as string
func (b Boolean) String() string {
	if b {
		return "true"
	}

	return "false"
}
