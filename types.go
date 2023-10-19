package icingadsl

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var indentation int

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

	bla.WriteString("object CheckCommand \"" + cc.Name + "\" {\n")

	indentation++

	for _, cci := range cc.Imports {
		bla.WriteString(IndntStr() + "import \"" + cci.Name + "\"\n")
	}

	bla.WriteString(IndntStr() + "command = " + cc.Command.String() + "\n")

	bla.WriteString(IndntStr() + "arguments = {\n")
	indentation++

	for i := range cc.Arguments {
		bla.WriteString(cc.Arguments[i].String(cc.Name))
		bla.WriteString("\n")
	}

	indentation--

	bla.WriteString(IndntStr() + "}\n")

	indentation--

	bla.WriteString(IndntStr() + "}\n")

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

	b.WriteString(IndntStr() + "\"" + cca.Name + "\" = {\n")

	indentation++

	if cca.Value != "" {
		if prefix != "" {
			b.WriteString(IndntStr() + "value = \"$" + prefix + "_" + strings.ReplaceAll(cca.Value, "-", "_") + "$\"\n")
		} else {
			b.WriteString(IndntStr() + "value = \"$" + strings.ReplaceAll(cca.Value, "-", "_") + "$\"\n")
		}
	} else {
		b.WriteString(IndntStr() + "value = \"\"\n")
	}

	if cca.Description != "" {
		b.WriteString(IndntStr() + "description = " + cca.Description.String() + "\n")
	}

	if cca.Required {
		b.WriteString(IndntStr() + "required = true\n")
	}

	if cca.SkipKey {
		b.WriteString(IndntStr() + "skip_key = true\n")
	}

	if cca.SetIf != nil {
		switch tmp := cca.SetIf.(type) {
		case String:
			if prefix != "" {
				b.WriteString(IndntStr() + "set_if = \"$" + prefix + "_" + strings.ReplaceAll(tmp.RawString(), "-", "_") + "$\"\n")
			} else {
				b.WriteString(IndntStr() + "set_if = \"$" + strings.ReplaceAll(tmp.RawString(), "-", "_") + "$\"\n")
			}
		case Boolean:
			b.WriteString(IndntStr() + "set_if = " + tmp.String() + "\n")
		default:
		}
	}

	if cca.Order != 0 {
		b.WriteString(IndntStr() + "order = " + fmt.Sprintf("%d", cca.Order) + "\n")
	}

	if !cca.RepeatKey {
		b.WriteString(IndntStr() + "repeat_key = false\n")
	}

	if cca.Key != "" {
		b.WriteString(IndntStr() + "key = \"" + cca.Key + "\"\n")
	}

	if cca.Separator != "" {
		b.WriteString(IndntStr() + "separator = \"" + cca.Separator + "\"\n")
	}

	indentation--

	b.WriteString(IndntStr() + "}")

	return b.String()
}

func IndntStr() string {
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
