package icingadsl

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCheckCommandArgument(t *testing.T) {
	cca := CheckCommandArgument{
		Name:        "--foo",
		Value:       "bla_foo",
		Description: "hello",
		Required:    false,
		SetIf:       String("bla_foo_bool"),
		Order:       5,
	}

	resultString := `"--foo" = {
	value = "$bla_foo$"
	description = "hello"
	set_if = "$bla_foo_bool$"
	order = 5
	repeat_key = false
}`

	assert.Equal(t, resultString, cca.String(""))
}

func TestString(t *testing.T) {
	string1 := String("foo\nbar")

	assert.Equal(t, "{{{foo\nbar}}}", string1.String())

}

func TestArray(t *testing.T) {
	ia := Array{
		String("foo"),
		String("bla"),
	}

	resultString := `["foo", "bla"]`

	assert.Equal(t, resultString, ia.String())

	ia2 := Array{}

	assert.Equal(t, "[]", ia2.String())

	ia3 := Array{
		String("foo"),
	}

	assert.Equal(t, "[\"foo\"]", ia3.String())
}

func TestCheckCommand(t *testing.T) {
	cc := CheckCommand{
		Name:    "MyPlugin",
		Command: Array{Identifier("PluginContribDir"), String("check_myPlugin")},
		Vars:    Dictionary{Identifier("var1"): Integer(56)},
		Timeout: Duration{time.Duration(30 * time.Second)},
		Arguments: []CheckCommandArgument{
			{
				Name:        "--foo",
				Value:       "foo_val",
				Description: String("hello\nneighbour"),
				Required:    false,
				SetIf:       String("bla_foo_bool"),
				Order:       5,
			},
			{
				Name:        "--bla",
				Value:       "bla_val",
				Description: "ciao",
				Required:    true,
				SetIf:       String("bla_foo_bool"),
			},
		},
	}

	resultString := `object CheckCommand "MyPlugin" {
	command = [PluginContribDir, "check_myPlugin"]
	arguments = {
		"--foo" = {
			value = "$MyPlugin_foo_val$"
			description = {{{hello
neighbour}}}
			set_if = "$MyPlugin_bla_foo_bool$"
			order = 5
			repeat_key = false
		}
		"--bla" = {
			value = "$MyPlugin_bla_val$"
			description = "ciao"
			required = true
			set_if = "$MyPlugin_bla_foo_bool$"
			repeat_key = false
		}
	}
}
`

	assert.Equal(t, resultString, cc.String())
}

func TestCheckCommandWithEmptyArgs(t *testing.T) {
	cc := CheckCommand{
		Name: "MyPlugin",
	}

	resultString := `object CheckCommand "MyPlugin" {
	command = []
	arguments = {
	}
}
`

	assert.Equal(t, resultString, cc.String())
}

func TestCheckCommandWithFilledArgs(t *testing.T) {
	cc := CheckCommand{
		Name:    "MyPlugin",
		Command: Array{Identifier("PluginContribDir"), String("check_myPlugin")},
		Vars:    Dictionary{Identifier("var1"): Integer(56)},
		Timeout: Duration{time.Duration(30 * time.Second)},
		Arguments: []CheckCommandArgument{
			{
				Name:        "--foo",
				Value:       "foo_val",
				Description: String("hello\"neighbour\""),
				Required:    false,
				SetIf:       String("bla_foo_bool"),
				Order:       5,
			},
			{
				Name:        "--bla",
				Value:       "bla_val",
				Description: "ciao",
				Required:    true,
				SetIf:       True,
			},
		},
	}

	resultString := `object CheckCommand "MyPlugin" {
	command = [PluginContribDir, "check_myPlugin"]
	arguments = {
		"--foo" = {
			value = "$MyPlugin_foo_val$"
			description = "hello\"neighbour\""
			set_if = "$MyPlugin_bla_foo_bool$"
			order = 5
			repeat_key = false
		}
		"--bla" = {
			value = "$MyPlugin_bla_val$"
			description = "ciao"
			required = true
			set_if = true
			repeat_key = false
		}
	}
}
`

	assert.Equal(t, resultString, cc.String())
}
