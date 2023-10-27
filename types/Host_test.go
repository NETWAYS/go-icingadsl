package icingadsl

import (
	"testing"
	"time"
)

func TestNullHost(t *testing.T) {
	host := Host{}
	host.Name = "foo"

	stringed_host := host.String()

	compareString := `object Host "foo" {
}`

	assertEqualString(t, stringed_host, compareString)
}

func TestFullHost(t *testing.T) {
	host := Host{}

	host.Name = "foo"

	htGeneric := HostTemplate{
		Name: "generic",
	}

	htSpecial := HostTemplate{
		Name: "special",
	}

	host.Imports = []*HostTemplate{
		&htGeneric,
		&htSpecial,
	}

	host.DisplayName = String("bla")

	host.Address = "192.0.2.4"
	host.Address6 = "2001:DB8::666"

	exampleGroup := HostGroup{
		Name: "example",
	}

	host.Groups = []*HostGroup{
		&exampleGroup,
	}

	host.Vars = Dictionary{
		"parameter1": String("nope"),
		"parameter2": String("never"),
	}

	hostalive := CheckCommand{
		Name: "hostalive",
	}

	host.CheckCommand = &hostalive

	host.maxCheckAttemptsSet = true
	host.MaxCheckAttempts = 5

	weirdTimePerioud := TimePeriod{
		Name: "EverySecondFullMoon",
	}

	host.CheckPeriod = &weirdTimePerioud

	tmpDur, err := time.ParseDuration("5m")

	if err != nil {
		t.Errorf("Error during Parsing \"5m\" as a duration")
	}

	host.CheckTimeout = Duration{
		duration: tmpDur,
	}
	host.checkTimeoutSet = true

	host.CheckInterval = Duration{
		duration: tmpDur,
	}
	host.checkIntervalSet = true

	host.RetryInterval = Duration{
		duration: tmpDur,
	}
	host.retryIntervalSet = true

	host.EnableNotifications = false
	host.enableNotificationsSet = true

	host.EnableActiveChecks = true
	host.enableActiveChecksSet = true

	host.EnablePassiveChecks = true
	host.enablePassiveChecksSet = true

	host.EnableEventHandler = false
	host.enableEventHandlerSet = true

	reboot := EventCommand{
		Name: "reboot",
	}

	host.EventCommand = &reboot

	host.FlappingThresholdLow = 1.0
	host.flappingThresholdLowSet = true

	host.FlappingThresholdHigh = 2.0
	host.flappingThresholdHighSet = true

	host.FlappingIgnoreStates = Array{
		String("WARNING"),
	}

	host.Volatile = true
	host.volatileSet = true

	main := Zone{
		Name: "main",
	}

	host.Zone = &main

	testShredder := Endpoint{
		Name: "testShredder",
	}

	host.CommandEndpoint = &testShredder

	host.Notes = String("Note to me:")
	host.NotesURL = String("https://icinga.example.com/docu")

	host.ActionURL = String("https://icinga.example.com/action/$host.name")

	host.IconImage = "exampleHost.png"
	host.IconImageAlt = "brokenExampleHost.png"

	stringed_host := host.String()

	compareString := `object Host "foo" {
	import "generic"
	import "special"
	display_name = "bla"
	address = "192.0.2.4"
	address6 = "2001:DB8::666"
	groups = ["example"]
	vars += {
		parameter1 = "nope",
		parameter2 = "never",
	}
	check_command = "hostalive"
	max_check_attempts = 5
	check_period = "EverySecondFullMoon"
	check_timeout = 300s
	check_interval = 300s
	retry_interval = 300s
	enable_notifications = false
	enable_active_checks = true
	enable_passive_checks = true
	enable_event_handler = false
	event_command = "reboot"
	flapping_threshold_high = 2.0
	flapping_threshold_low = 1.0
	flapping_ignore_states = ["WARNING"]
	volatile = true
	zone = "main"
	command_endpoint = "testShredder"
	notes = "Note to me:"
	notes_url = "https://icinga.example.com/docu"
	action_url = "https://icinga.example.com/action/$host.name"
	icon_image = "exampleHost.png"
	icon_image_alt = "brokenExampleHost.png"
}`

	assertEqualString(t, stringed_host, compareString)
}
