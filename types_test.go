package icingadsl

import (
	"testing"
	"time"
)

func assertEqualString(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Error("\nActual: ", actual, "\nExpected: ", expected)
	}
}

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

	assertEqualString(t, resultString, cca.String(""))
}

func TestInfixExpression(t *testing.T) {
	int1 := InfixExpression{
		Left:          Boolean(true),
		InfixOperator: PLUS,
		Right:         Boolean(false),
	}

	assertEqualString(t, "true + false", int1.String())
}

func TestBoolean(t *testing.T) {
	b1 := Boolean(true)
	b2 := Boolean(false)

	assertEqualString(t, "true", b1.String())
	assertEqualString(t, "false", b2.String())
}

func TestInteger(t *testing.T) {
	int1 := Integer(1)

	assertEqualString(t, "1", int1.String())
}

func TestString(t *testing.T) {
	string1 := String("foo\nbar")

	assertEqualString(t, "{{{foo\nbar}}}", string1.String())
}

func TestArray(t *testing.T) {
	ia := Array{
		String("foo"),
		String("bla"),
	}

	resultString := `["foo", "bla"]`

	assertEqualString(t, resultString, ia.String())

	ia2 := Array{}

	assertEqualString(t, "[]", ia2.String())

	ia3 := Array{
		String("foo"),
	}

	assertEqualString(t, "[\"foo\"]", ia3.String())
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

	assertEqualString(t, resultString, cc.String())
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

	assertEqualString(t, resultString, cc.String())
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

	assertEqualString(t, resultString, cc.String())
}
