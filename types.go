package icingadsl

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Duration struct {
	time.Duration
}

type Integer int

type String string

type Identifier string

type Dictionary map[Identifier]Object

type Array []Object

type Boolean bool

type Object interface {
	String() string
}

const (
	True  = Boolean(true)
	False = Boolean(false)
)

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

	var indentation int

	bla.WriteString("object CheckCommand \"" + cc.Name + "\" {\n")

	indentation++

	for _, cci := range cc.Imports {
		bla.WriteString(indentString(indentation) + "import \"" + cci.Name + "\"\n")
	}

	bla.WriteString(indentString(indentation) + "command = " + cc.Command.String() + "\n")

	bla.WriteString(indentString(indentation) + "arguments = {\n")
	indentation++

	for i := range cc.Arguments {
		bla.WriteString(cc.Arguments[i].String(indentation, cc.Name))
		bla.WriteString("\n")
	}

	indentation--

	bla.WriteString(indentString(indentation) + "}\n")

	indentation--

	bla.WriteString(indentString(indentation) + "}\n")

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

func (cca *CheckCommandArgument) String(indentation int, prefix string) string {
	var b strings.Builder

	b.WriteString(indentString(indentation) + "\"" + cca.Name + "\" = {\n")

	indentation++

	if cca.Value != "" {
		if prefix != "" {
			b.WriteString(indentString(indentation) + "value = \"$" + prefix + "_" + strings.ReplaceAll(cca.Value, "-", "_") + "$\"\n")
		} else {
			b.WriteString(indentString(indentation) + "value = \"$" + strings.ReplaceAll(cca.Value, "-", "_") + "$\"\n")
		}
	} else {
		b.WriteString(indentString(indentation) + "value = \"\"\n")
	}

	if cca.Description != "" {
		b.WriteString(indentString(indentation) + "description = " + cca.Description.String() + "\n")
	}

	if cca.Required {
		b.WriteString(indentString(indentation) + "required = true\n")
	}

	if cca.SkipKey {
		b.WriteString(indentString(indentation) + "skip_key = true\n")
	}

	if cca.SetIf != nil {
		switch tmp := cca.SetIf.(type) {
		case String:
			if prefix != "" {
				b.WriteString(indentString(indentation) + "set_if = \"$" + prefix + "_" + strings.ReplaceAll(tmp.RawString(), "-", "_") + "$\"\n")
			} else {
				b.WriteString(indentString(indentation) + "set_if = \"$" + strings.ReplaceAll(tmp.RawString(), "-", "_") + "$\"\n")
			}
		case Boolean:
			b.WriteString(indentString(indentation) + "set_if = " + tmp.String() + "\n")
		default:
		}
	}

	if cca.Order != 0 {
		b.WriteString(indentString(indentation) + "order = " + fmt.Sprintf("%d", cca.Order) + "\n")
	}

	if !cca.RepeatKey {
		b.WriteString(indentString(indentation) + "repeat_key = false\n")
	}

	if cca.Key != "" {
		b.WriteString(indentString(indentation) + "key = \"" + cca.Key + "\"\n")
	}

	if cca.Separator != "" {
		b.WriteString(indentString(indentation) + "separator = \"" + cca.Separator + "\"\n")
	}

	indentation--

	b.WriteString(indentString(indentation) + "}")

	return b.String()
}

func indentString(i int) string {
	return strings.Repeat("\t", i)
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
