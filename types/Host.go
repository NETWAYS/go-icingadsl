package icingadsl

import (
	"strconv"
	"strings"
)

// https://icinga.com/docs/icinga-2/latest/doc/09-object-types/#host
type Host struct {
	Name string

	// Imports is, strictly speaking, broken by definition here since imports
	// will overwrite values depending on wheter they are above or below other directives.
	// Therefore this covers only the case "Imported before any specific host configuration"
	Imports []*HostTemplate `icingadsl:"import"`

	DisplayName String     `icingadsl:"display_name"`
	Address     String     `icingadsl:"address"`
	Address6    String     `icingadsl:"address6"`
	Groups      Array      `icingadsl:"groups"`
	Vars        Dictionary `icingadsl:"vars"`

	CheckCommand *CheckCommand `icingadsl:"check_command"`

	MaxCheckAttempts    uint64 `icingadsl:"max_check_attempts"`
	maxCheckAttemptsSet bool

	CheckPeriod *TimePeriod `icingadsl:"check_period"`

	CheckTimeout    Duration `icingadsl:"check_timeout"`
	checkTimeoutSet bool

	CheckInterval    Duration `icingadsl:"check_interval"`
	checkIntervalSet bool

	RetryInterval    Duration `icingadsl:"retry_interval"`
	retryIntervalSet bool

	EnableNotifications    Boolean `icingadsl:"enable_notifications"`
	enableNotificationsSet bool

	EnableActiveChecks    Boolean `icingadsl:"enable_active_checks"`
	enableActiveChecksSet bool

	EnablePassiveChecks    Boolean `icingadsl:"enable_passive_checks"`
	enablePassiveChecksSet bool

	EnableEventHandler    Boolean `icingadsl:"enable_event_handler"`
	enableEventHandlerSet bool

	EventCommand *EventCommand `icingadsl:"event_command"`

	FlappingThresholdHigh    Float `icingadsl:"flapping_threshold_high"`
	flappingThresholdHighSet bool

	FlappingThresholdLow    Float `icingadsl:"flapping_threshold_low"`
	flappingThresholdLowSet bool

	FlappingIgnoreStates Array `icingadsl:"flapping_ignore_states"`

	Volatile    Boolean `icingadsl:"voltile"`
	volatileSet bool

	Zone *Zone `icingadsl:"zone"`

	CommandEndpoint *Endpoint `icingadsl:"command_endpoint"`

	Notes    String `icingadsl:"notes"`
	NotesUrl String `icingadsl:"notes_url"`

	ActionUrl String `icingadsl:"action_url"`

	IconImage    String `icingadsl:"icon_image"`
	IconImageAlt String `icingadsl:"icon_image_alt"`
}

func (h *Host) String() string {
	var stringer strings.Builder

	stringer.WriteString("object Host \"" + h.Name + "\" {\n")

	indentation++

	if len(h.Imports) != 0 {
		for index := range h.Imports {
			stringer.WriteString(indentString() + "import \"" + h.Imports[index].GetName() + "\"\n")
		}
	}

	if h.DisplayName != "" {
		stringer.WriteString(indentString() + "display_name = " + h.DisplayName.String() + "\n")
	}

	if h.Address != "" {
		stringer.WriteString(indentString() + "address = " + h.Address.String() + "\n")
	}

	if h.Address6 != "" {
		stringer.WriteString(indentString() + "address6 = " + h.Address6.String() + "\n")
	}

	if len(h.Groups) != 0 {
		stringer.WriteString(indentString() + "groups = " + h.Groups.String() + "\n")
	}

	if len(h.Vars) != 0 {
		stringer.WriteString(indentString() + "vars += " + h.Vars.String() + "\n")
	}

	if h.CheckCommand != nil {
		stringer.WriteString(indentString() + "check_command = \"" + h.CheckCommand.Name + "\"\n")
	}

	if h.maxCheckAttemptsSet {
		stringer.WriteString(indentString() + "max_check_attempts = " + strconv.FormatUint(h.MaxCheckAttempts, 10) + "\n")
	}

	if h.CheckPeriod != nil {
		stringer.WriteString(indentString() + "check_period = \"" + h.CheckPeriod.Name + "\"\n")
	}

	if h.checkTimeoutSet {
		stringer.WriteString(indentString() + "check_timeout = " + h.CheckTimeout.String() + "\n")
	}

	if h.checkIntervalSet {
		stringer.WriteString(indentString() + "check_interval = " + h.CheckInterval.String() + "\n")
	}

	if h.retryIntervalSet {
		stringer.WriteString(indentString() + "retry_interval = " + h.RetryInterval.String() + "\n")
	}

	if h.enableNotificationsSet {
		stringer.WriteString(indentString() + "enable_notifications = " + h.EnableNotifications.String() + "\n")
	}

	if h.enableActiveChecksSet {
		stringer.WriteString(indentString() + "enable_active_checks = " + h.EnableActiveChecks.String() + "\n")
	}

	if h.enablePassiveChecksSet {
		stringer.WriteString(indentString() + "enable_passive_checks = " + h.EnablePassiveChecks.String() + "\n")
	}

	if h.enableEventHandlerSet {
		stringer.WriteString(indentString() + "enable_event_handler = " + h.EnableEventHandler.String() + "\n")
	}

	if h.EventCommand != nil {
		stringer.WriteString(indentString() + "event_command = \"" + h.EventCommand.Name + "\"\n")
	}

	if h.flappingThresholdHighSet {
		stringer.WriteString(indentString() + "flapping_threshold_high = " + h.FlappingThresholdHigh.String() + "\n")
	}

	if h.flappingThresholdLowSet {
		stringer.WriteString(indentString() + "flapping_threshold_low = " + h.FlappingThresholdLow.String() + "\n")
	}

	if len(h.FlappingIgnoreStates) != 0 {
		stringer.WriteString(indentString() + "flapping_ignore_states = " + h.FlappingIgnoreStates.String() + "\n")
	}

	if h.volatileSet {
		stringer.WriteString(indentString() + "volatile = " + h.Volatile.String() + "\n")
	}

	if h.Zone != nil {
		stringer.WriteString(indentString() + "zone = \"" + h.Zone.Name + "\"\n")
	}

	if h.CommandEndpoint != nil {
		stringer.WriteString(indentString() + "command_endpoint = \"" + h.CommandEndpoint.Name + "\"\n")
	}

	if h.Notes != "" {
		stringer.WriteString(indentString() + "notes = " + h.Notes.String() + "\n")
	}

	if h.NotesUrl != "" {
		stringer.WriteString(indentString() + "notes_url = " + h.NotesUrl.String() + "\n")
	}

	if h.ActionUrl != "" {
		stringer.WriteString(indentString() + "action_url = " + h.ActionUrl.String() + "\n")
	}

	if h.IconImage != "" {
		stringer.WriteString(indentString() + "icon_image = " + h.IconImage.String() + "\n")
	}

	if h.IconImageAlt != "" {
		stringer.WriteString(indentString() + "icon_image_alt = " + h.IconImageAlt.String() + "\n")
	}

	indentation--
	stringer.WriteString(indentString() + "}")

	return stringer.String()
}
