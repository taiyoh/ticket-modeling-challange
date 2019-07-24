package domain

import "time"

// TimeWindow provides price segment by supplied date.
type TimeWindow struct {
	t       time.Time
	holiday bool
}

// TimeWindowType represents segment types for price by time.
type TimeWindowType int

const (
	// WeekDayDayTime represents TimeWindowType which explains weekday and daytime.
	WeekDayDayTime TimeWindowType = iota + 1
	// WeekDayNightTime represents TimeWindowType which explains weekday and night-time.
	WeekDayNightTime
	// HolidayDayTime represents TimeWindowType which explains holiday and daytime.
	HolidayDayTime
	// HolidayNightTime represents TimeWindowType which explains holiday and night-time.
	HolidayNightTime
	// MovieDay represents TimeWindowType which explain the first day of this month.
	MovieDay
)

type dayType int

const (
	weekday dayType = iota + 1
	holiday
)

type hourType int

const (
	daytime hourType = iota + 1
	nighttime
)

var typeByDayAndHour = map[dayType]map[hourType]TimeWindowType{
	weekday: map[hourType]TimeWindowType{
		daytime:   WeekDayDayTime,
		nighttime: WeekDayNightTime,
	},
	holiday: map[hourType]TimeWindowType{
		daytime:   HolidayDayTime,
		nighttime: HolidayNightTime,
	},
}

func (w TimeWindow) dayType() dayType {
	wd := w.t.Weekday()
	if w.holiday || wd == time.Saturday || wd == time.Sunday {
		return holiday
	}
	return weekday
}

func (w TimeWindow) hourType() hourType {
	if w.t.Hour() < 20 {
		return daytime
	}
	return nighttime
}

// Type calculates TimeWindowType by supplied time.
func (w TimeWindow) Type() TimeWindowType {
	if w.t.Day() == 1 {
		return MovieDay
	}
	t, _ := typeByDayAndHour[w.dayType()][w.hourType()]
	return t
}
