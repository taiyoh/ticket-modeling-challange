package domain

import "time"

func NewTimeWindow(t time.Time, isHoliday bool) TimeWindow {
	return TimeWindow{t, isHoliday}
}
