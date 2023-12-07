package types

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
func FormatNotificationType(nt NotificationType) (string, error) {
	switch nt {
	case DowntimeStart:
		return "downtimestart", nil
	case DowntimeEnd:
		return "downtimeend", nil
	case DowntimeRemoved:
		return "downtimeremoved", nil
	case Custom:
		return "custom", nil
	case Acknowledgement:
		return "acknowledgement", nil
	case Problem:
		return "problem", nil
	case Recovery:
		return "recovery", nil
	case FlappingStart:
		return "flappingstart", nil
	case FlappingEnd:
		return "flappingend", nil
	default:
		return "", errors.New("no matching state for the provided number")
	}
}
