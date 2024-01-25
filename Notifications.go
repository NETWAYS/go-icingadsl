package icingadsl

import (
	"errors"
	"strings"
)

type NotificationType uint

const (
	DowntimeStart NotificationType = iota
	DowntimeEnd
	DowntimeRemoved
	Custom
	Acknowledgement
	Problem
	Recovery
	FlappingStart
	FlappingEnd
)

/*
 * Parse a notification type string (typically received when Icinga 2 is executing a notification plugin)
 * into a fitting enum to simplify the following logic.
 */
func ParseNotificationType(nt string) (NotificationType, error) {
	switch strings.ToLower(nt) {
	case "downtimestart":
		return DowntimeStart, nil
	case "downtimeend":
		return DowntimeEnd, nil
	case "downtimeremoved":
		return DowntimeRemoved, nil
	case "custom":
		return Custom, nil
	case "acknowledgement":
		return Acknowledgement, nil
	case "problem":
		return Problem, nil
	case "recovery":
		return Recovery, nil
	case "flappingstart":
		return FlappingStart, nil
	case "flappingend":
		return FlappingEnd, nil
	default:
		return 0, errors.New("no matching state for the provided string")
	}
}

/*
 * Transforms a notification type into a string
 */
func (nt NotificationType) String() string {
	switch nt {
	case DowntimeStart:
		return "downtimestart"
	case DowntimeEnd:
		return "downtimeend"
	case DowntimeRemoved:
		return "downtimeremoved"
	case Custom:
		return "custom"
	case Acknowledgement:
		return "acknowledgement"
	case Problem:
		return "problem"
	case Recovery:
		return "recovery"
	case FlappingStart:
		return "flappingstart"
	case FlappingEnd:
		return "flappingend"
	default:
		return ""
	}
}
