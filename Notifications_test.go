package icingadsl

import (
	"testing"
)

func TestParseNotificationType(t *testing.T) {
	testcases := []struct {
		expected NotificationType
		nt       string
	}{
		{
			expected: DowntimeStart,
			nt:       "downtimestart",
		},
		{
			expected: DowntimeEnd,
			nt:       "downtimeend",
		},
		{
			expected: DowntimeRemoved,
			nt:       "downtimeremoved",
		},
		{
			expected: Custom,
			nt:       "custom",
		},
		{
			expected: Acknowledgement,
			nt:       "acknowledgement",
		},
		{
			expected: Problem,
			nt:       "problem",
		},
		{
			expected: Recovery,
			nt:       "recovery",
		},
		{
			expected: FlappingStart,
			nt:       "flappingstart",
		},
		{
			expected: FlappingEnd,
			nt:       "flappingend",
		},
	}

	for _, tc := range testcases {
		nt, err := ParseNotificationType(tc.nt)

		if err != nil {
			t.Error("Did not expect error, got:", err)
		}

		if nt != tc.expected {
			t.Error("\nActual: ", nt, "\nExpected: ", tc.expected)
		}
	}
}

func TestNotificationTypeStringer(t *testing.T) {
	testcases := []struct {
		nt       NotificationType
		expected string
	}{
		{
			nt:       DowntimeStart,
			expected: "downtimestart",
		},
		{
			nt:       DowntimeEnd,
			expected: "downtimeend",
		},
		{
			nt:       DowntimeRemoved,
			expected: "downtimeremoved",
		},
		{
			nt:       Custom,
			expected: "custom",
		},
		{
			nt:       Acknowledgement,
			expected: "acknowledgement",
		},
		{
			nt:       Problem,
			expected: "problem",
		},
		{
			nt:       Recovery,
			expected: "recovery",
		},
		{
			nt:       FlappingStart,
			expected: "flappingstart",
		},
		{
			nt:       FlappingEnd,
			expected: "flappingend",
		},
	}

	for _, tc := range testcases {
		if tc.nt.String() != tc.expected {
			t.Error("\nActual: ", tc.nt.String(), "\nExpected: ", tc.expected)
		}
	}
}
