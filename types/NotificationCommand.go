package icingadsl

import (
	"strings"
)

// https://icinga.com/docs/icinga-2/latest/doc/09-object-types/#notificationcommand
type NotificationCommand struct {
	Name      string
	Command   Array
	Imports   []*NotificationCommand
	Env       map[string]string
	Vars      Dictionary
	Timeout   Duration
	Arguments []CommandArgument
}

// String returns the CheckCommand Object as string with proper indentation
func (cc *NotificationCommand) String() string {
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
