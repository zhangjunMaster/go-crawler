package scheduler

type Status uint8

const (
	SCHED_STATUS_UNINITALIZED Status = 0
	SCHED_STATUS_INITALIZING  Status = 1
	SCHED_STATUS_INITALIZED   Status = 2
	SCHED_STATUS_STARTING     Status = 3
	SCHED_STATUS_STARTED      Status = 4
	SCHED_STATUS_STOPING      Status = 5
	SCHED_STATUS_STOPPED      Status = 6
)
