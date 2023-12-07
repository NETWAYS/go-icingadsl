package types

import (
	"fmt"
	"strings"
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

// String returns the CheckCommandArgument Object as string with the given prefix
// and proper indentation
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
