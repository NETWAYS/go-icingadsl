package icingadsl

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var indentation int

type Identifier string

// https://icinga.com/docs/icinga-2/latest/doc/17-language-reference/#duration-literals
type Duration struct {
	time.Duration
}

// https://icinga.com/docs/icinga-2/latest/doc/18-library-reference/#number-type
// This is only the integer "half" of the Icinga2 type
type Integer int

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
// TODO

const (
	True  = Boolean(true)
	False = Boolean(false)
)

// https://icinga.com/docs/icinga-2/latest/doc/09-object-types/#checkcommand-arguments
type CheckCommandArgument struct {
	Name        string
	Value       string
	Description String
	SetIf       Object
	Separator   string
	Key         string
	Order       int
	RepeatKey   bool
	Required    bool
	SkipKey     bool
}

// https://icinga.com/docs/icinga-2/latest/doc/09-object-types/#checkcommand
type CheckCommand struct {
	Name      string
	Command   Array
	Imports   []*CheckCommand
	Env       map[string]string
	Vars      Dictionary
	Timeout   Duration
	Arguments []CheckCommandArgument
}

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

func (is String) RawString() string {
	return string(is)
}

func (i Identifier) String() string {
	return string(i)
}

func (cc *CheckCommand) String() string {
	var bla strings.Builder

	bla.WriteString("object CheckCommand \"" + cc.Name + "\" {\n")

	indentation++

	for _, cci := range cc.Imports {
		bla.WriteString(indentString() + "import \"" + cci.Name + "\"\n")
	}

	bla.WriteString(indentString() + "command = " + cc.Command.String() + "\n")

	bla.WriteString(indentString() + "arguments = {\n")
	indentation++

	for i := range cc.Arguments {
		bla.WriteString(cc.Arguments[i].String(cc.Name))
		bla.WriteString("\n")
	}

	indentation--

	bla.WriteString(indentString() + "}\n")

	indentation--

	bla.WriteString(indentString() + "}\n")

	return bla.String()
}

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

func (cca *CheckCommandArgument) String(prefix string) string {
	var b strings.Builder

	b.WriteString(indentString() + "\"" + cca.Name + "\" = {\n")

	indentation++

	if cca.Value != "" {
		if prefix != "" {
			b.WriteString(indentString() + "value = \"$" + prefix + "_" + strings.ReplaceAll(cca.Value, "-", "_") + "$\"\n")
		} else {
			b.WriteString(indentString() + "value = \"$" + strings.ReplaceAll(cca.Value, "-", "_") + "$\"\n")
		}
	} else {
		b.WriteString(indentString() + "value = \"\"\n")
	}

	if cca.Description != "" {
		b.WriteString(indentString() + "description = " + cca.Description.String() + "\n")
	}

	if cca.Required {
		b.WriteString(indentString() + "required = true\n")
	}

	if cca.SkipKey {
		b.WriteString(indentString() + "skip_key = true\n")
	}

	if cca.SetIf != nil {
		switch tmp := cca.SetIf.(type) {
		case String:
			if prefix != "" {
				b.WriteString(indentString() + "set_if = \"$" + prefix + "_" + strings.ReplaceAll(tmp.RawString(), "-", "_") + "$\"\n")
			} else {
				b.WriteString(indentString() + "set_if = \"$" + strings.ReplaceAll(tmp.RawString(), "-", "_") + "$\"\n")
			}
		case Boolean:
			b.WriteString(indentString() + "set_if = " + tmp.String() + "\n")
		default:
		}
	}

	if cca.Order != 0 {
		b.WriteString(indentString() + "order = " + fmt.Sprintf("%d", cca.Order) + "\n")
	}

	if !cca.RepeatKey {
		b.WriteString(indentString() + "repeat_key = false\n")
	}

	if cca.Key != "" {
		b.WriteString(indentString() + "key = \"" + cca.Key + "\"\n")
	}

	if cca.Separator != "" {
		b.WriteString(indentString() + "separator = \"" + cca.Separator + "\"\n")
	}

	indentation--

	b.WriteString(indentString() + "}")

	return b.String()
}

func indentString() string {
	return strings.Repeat("\t", indentation)
}

func (op *Operator) String() string {
	if *op == PLUS {
		return "+"
	}

	return ""
}

func (ie InfixExpression) String() string {
	var result strings.Builder

	result.WriteString(ie.Left.String() + " ")
	result.WriteString(ie.InfixOperator.String() + " ")
	result.WriteString((ie.Right).String())

	return result.String()
}

func (b Boolean) String() string {
	if b {
		return "true"
	}

	return "false"
}
