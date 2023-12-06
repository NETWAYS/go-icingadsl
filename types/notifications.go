package icingadsl

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
