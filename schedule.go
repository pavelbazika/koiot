package koiot

import "time"

type (
	ScheduleTime struct {
		Year    *int          `yaml:"year"`
		Month   *time.Month   `yaml:"month"`
		Day     *int          `yaml:"day"`
		WeekDay *time.Weekday `yaml:"weekday"`
		Hour    *int          `yaml:"hour"`
		Minute  *int          `yaml:"minute"`
		Second  *int          `yaml:"second"`
	}

	ScheduleItem[TData any] struct {
		When ScheduleTime `yaml:"when"`
		What TData        `yaml:"what"`
	}

	Schedule[TData any] []ScheduleItem[TData]
)

func (schedule Schedule[TData]) GetCurrentItemData(t time.Time) any {
	// find the nearest record
	winner := -1
	var winnerTime time.Time
	for i, sch := range schedule {
		trigTime, found := sch.When.GetPrevTime(t)
		if !found {
			continue
		}
		if trigTime.Equal(t) || (trigTime.Before(t) && (winnerTime.IsZero() || trigTime.After(winnerTime))) {
			winnerTime = trigTime
			winner = i
		}
		if trigTime.Equal(t) {
			break
		}
	}

	if winner != -1 {
		return schedule[winner].What
	} else {
		return nil
	}
}

// get time when the item triggered last before (or at) t
func (sch *ScheduleTime) GetPrevTime(t time.Time) (time.Time, bool) {
	if (sch.WeekDay != nil) && (sch.Day != nil) {
		panic("Cannot set both day and weekday")
	}

	// extract given time
	rqHr, rqMin, rqSec := t.Clock()
	rqYr, rqMon, rqDay := t.Date()

	var yr int
	var mon time.Month
	var day int
	var hr int
	var min int
	var sec int

	prevTime := false

	// year
	if sch.Year != nil {
		yr = *sch.Year
	} else {
		yr = rqYr
	}
	if yr > rqYr {
		// first ocurrence is in the future
		return time.Time{}, false
	} else if yr < rqYr {
		// choosen year is in the past
		prevTime = true
	}

	// month
	switch {
	case sch.Month != nil:
		mon = *sch.Month

	case prevTime:
		mon = 12

	default:
		mon = rqMon
	}
	if !prevTime && (mon > rqMon) {
		// month specified by schedule, after rq
		if sch.Year == nil {
			// year can be manipulated
			yr--
		} else {
			// year cannot be moved, event in future
			return time.Time{}, false
		}
	}
	prevTime = (yr < rqYr) || (mon < rqMon)

	// day
	switch {
	case sch.Day != nil:
		day = *sch.Day

	case prevTime:
		day = daysIn(yr, mon)

	default:
		day = rqDay
	}
	if sch.WeekDay != nil {
		if !sch.subYMD(&yr, &mon, &day, true) {
			return time.Time{}, false
		}
		prevTime = (yr < rqYr) || (mon < rqMon) || (day < rqDay)
	}
	if !prevTime && (day > rqDay) {
		if !sch.subYM(&yr, &mon) {
			return time.Time{}, false
		}

		maxDay := daysIn(yr, mon)
		if maxDay < day {
			day = maxDay
		}
	}
	prevTime = (yr < rqYr) || (mon < rqMon) || (day < rqDay)

	switch {
	case sch.Hour != nil:
		hr = *sch.Hour

	case prevTime:
		hr = 23

	default:
		hr = rqHr
	}
	if !prevTime && (hr > rqHr) {
		if !sch.subYMD(&yr, &mon, &day, false) {
			return time.Time{}, false
		}
	}
	prevTime = (yr < rqYr) || (mon < rqMon) || (day < rqDay) || (hr < rqHr)

	switch {
	case sch.Minute != nil:
		min = *sch.Minute

	case prevTime:
		min = 59

	default:
		min = rqMin
	}
	if !prevTime && (min > rqMin) {
		if !sch.subYMDH(&yr, &mon, &day, &hr) {
			return time.Time{}, false
		}
	}
	prevTime = (yr < rqYr) || (mon < rqMon) || (day < rqDay) || (hr < rqHr) || (min < rqMin)

	switch {
	case sch.Second != nil:
		sec = *sch.Second

	case prevTime:
		sec = 59

	default:
		sec = rqSec
	}
	if !prevTime && (sec > rqSec) {
		if !sch.subYMDHM(&yr, &mon, &day, &hr, &min) {
			return time.Time{}, false
		}
	}

	return time.Date(yr, mon, day, hr, min, sec, 0, time.UTC), true
}

// get time when the item triggered last before (or at) t
func (sch *ScheduleTime) GetNextTime(t time.Time) (time.Time, bool) {
	if (sch.WeekDay != nil) && (sch.Day != nil) {
		panic("Cannot set both day and weekday")
	}

	// extract given time
	rqHr, rqMin, rqSec := t.Clock()
	rqYr, rqMon, rqDay := t.Date()

	var yr int
	var mon time.Month
	var day int
	var hr int
	var min int
	var sec int

	futureTime := false

	// year
	if sch.Year != nil {
		yr = *sch.Year
	} else {
		yr = rqYr
	}
	if yr < rqYr {
		// last ocurrence is in the past
		return time.Time{}, false
	} else if yr > rqYr {
		// choosen year is in the future
		futureTime = true
	}

	// month
	switch {
	case sch.Month != nil:
		mon = *sch.Month

	case futureTime:
		mon = 1

	default:
		mon = rqMon
	}
	if !futureTime && (mon < rqMon) {
		// month specified by schedule, before rq
		if sch.Year == nil {
			// year can be manipulated
			yr++
		} else {
			// year cannot be moved, event in past
			return time.Time{}, false
		}
	}
	futureTime = (yr > rqYr) || (mon > rqMon)

	// day
	switch {
	case sch.Day != nil:
		day = *sch.Day

	case futureTime:
		day = 1

	default:
		day = rqDay
	}
	if sch.WeekDay != nil {
		if !sch.addYMD(&yr, &mon, &day, true) {
			return time.Time{}, false
		}
		futureTime = (yr > rqYr) || (mon > rqMon) || (day > rqDay)
	}
	if !futureTime && (day < rqDay) {
		if !sch.addYM(&yr, &mon) {
			return time.Time{}, false
		}

		maxDay := daysIn(yr, mon)
		if maxDay < day {
			day = maxDay
		}
	}
	futureTime = (yr > rqYr) || (mon > rqMon) || (day > rqDay)

	switch {
	case sch.Hour != nil:
		hr = *sch.Hour

	case futureTime:
		hr = 0

	default:
		hr = rqHr
	}
	if !futureTime && (hr < rqHr) {
		if !sch.addYMD(&yr, &mon, &day, false) {
			return time.Time{}, false
		}
	}
	futureTime = (yr > rqYr) || (mon > rqMon) || (day > rqDay) || (hr > rqHr)

	switch {
	case sch.Minute != nil:
		min = *sch.Minute

	case futureTime:
		min = 0

	default:
		min = rqMin
	}
	if !futureTime && (min < rqMin) {
		if !sch.addYMDH(&yr, &mon, &day, &hr) {
			return time.Time{}, false
		}
	}
	futureTime = (yr > rqYr) || (mon > rqMon) || (day > rqDay) || (hr > rqHr) || (min > rqMin)

	switch {
	case sch.Second != nil:
		sec = *sch.Second

	case futureTime:
		sec = 0

	default:
		sec = rqSec
	}
	if !futureTime && (sec < rqSec) {
		if !sch.addYMDHM(&yr, &mon, &day, &hr, &min) {
			return time.Time{}, false
		}
	}

	return time.Date(yr, mon, day, hr, min, sec, 0, time.UTC), true
}

func daysIn(year int, m time.Month) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

// subtract month to previous ocurrence
func (sch *ScheduleTime) subYM(year *int, month *time.Month) bool {
	if sch.Month == nil {
		// month can be manipulated
		if *month == time.January {
			if sch.Year == nil {
				*year--
				*month = time.December
			} else {
				// year cannot be diminished
				return false
			}
		} else {
			*month--
		}
	} else {
		// month cannot be manipulated
		if sch.Year == nil {
			*year--
		} else {
			return false
		}
	}

	return true
}

func (sch *ScheduleTime) subYMD(year *int, month *time.Month, day *int, weekDayFind bool) bool {
	if sch.Day == nil {
		if sch.WeekDay == nil {
			// day can be freely manipulated
			if *day == 1 {
				if !sch.subYM(year, month) {
					return false
				}
				*day = daysIn(*year, *month)
			} else {
				*day--
			}
		} else {
			// weekday is set
			dayWd := time.Date(*year, *month, *day, 0, 0, 0, 0, time.UTC).Weekday()
			diff := 0
			if dayWd == *sch.WeekDay {
				// we are on requested weekday
				if !weekDayFind {
					diff = 7
				}
			} else {
				diff = int(dayWd) - int(*sch.WeekDay)
				if diff < 0 {
					diff += 7
				}
			}

			*day -= diff
			if *day < 1 {
				if !sch.subYM(year, month) {
					return false
				}
				*day = daysIn(*year, *month)
				dayWd = time.Date(*year, *month, *day, 0, 0, 0, 0, time.UTC).Weekday()
				if dayWd != *sch.WeekDay {
					diff = int(dayWd) - int(*sch.WeekDay)
					if diff < 0 {
						diff += 7
					}
				} else {
					diff = 0
				}

				*day -= diff
			}
		}
	} else {
		// day cannot be manipulated
		if !sch.subYM(year, month) {
			return false
		}

		// but when the day is after end of month, we manipulate
		maxDay := daysIn(*year, *month)
		if maxDay < *day {
			*day = maxDay
		}
	}

	return true
}

func (sch *ScheduleTime) subYMDH(year *int, month *time.Month, day *int, hour *int) bool {
	if sch.Hour == nil {
		// hour can be manipulated
		if *hour == 0 {
			if !sch.subYMD(year, month, day, false) {
				return false
			}
			*hour = 23
		} else {
			*hour--
		}

		return true
	} else {
		// static hour
		return sch.subYMD(year, month, day, false)
	}
}

func (sch *ScheduleTime) subYMDHM(year *int, month *time.Month, day *int, hour *int, min *int) bool {
	if sch.Minute == nil {
		// minute can be manipulated
		if *min == 0 {
			if !sch.subYMDH(year, month, day, hour) {
				return false
			}
			*min = 59
		} else {
			*min--
		}

		return true
	} else {
		return sch.subYMDH(year, month, day, hour)
	}
}

// add month to next ocurrence
func (sch *ScheduleTime) addYM(year *int, month *time.Month) bool {
	if sch.Month == nil {
		// month can be manipulated
		if *month == time.December {
			if sch.Year == nil {
				*year++
				*month = time.January
			} else {
				// year cannot be augmented
				return false
			}
		} else {
			*month++
		}
	} else {
		// month cannot be manipulated
		if sch.Year == nil {
			*year++
		} else {
			return false
		}
	}

	return true
}

func (sch *ScheduleTime) addYMD(year *int, month *time.Month, day *int, weekDayFind bool) bool {
	if sch.Day == nil {
		if sch.WeekDay == nil {
			// day can be freely manipulated
			if *day == daysIn(*year, *month) {
				if !sch.addYM(year, month) {
					return false
				}
				*day = 1
			} else {
				*day++
			}
		} else {
			// weekday is set
			dayWd := time.Date(*year, *month, *day, 0, 0, 0, 0, time.UTC).Weekday()
			diff := 0
			if dayWd == *sch.WeekDay {
				// we are on requested weekday
				if !weekDayFind {
					diff = 7
				}
			} else {
				diff = int(*sch.WeekDay) - int(dayWd)
				if diff < 0 {
					diff += 7
				}
			}

			*day += diff
			if *day > daysIn(*year, *month) {
				if !sch.addYM(year, month) {
					return false
				}
				*day = 1
				dayWd = time.Date(*year, *month, *day, 0, 0, 0, 0, time.UTC).Weekday()
				if dayWd != *sch.WeekDay {
					diff = int(*sch.WeekDay) - int(dayWd)
					if diff < 0 {
						diff += 7
					}
				} else {
					diff = 0
				}

				*day += diff
			}
		}
	} else {
		// day cannot be manipulated
		if !sch.addYM(year, month) {
			return false
		}

		// but when the day is after end of month, we manipulate
		maxDay := daysIn(*year, *month)
		if maxDay < *day {
			*day = maxDay
		}
	}

	return true
}

func (sch *ScheduleTime) addYMDH(year *int, month *time.Month, day *int, hour *int) bool {
	if sch.Hour == nil {
		// hour can be manipulated
		if *hour == 23 {
			if !sch.addYMD(year, month, day, false) {
				return false
			}
			*hour = 0
		} else {
			*hour++
		}

		return true
	} else {
		// static hour
		return sch.addYMD(year, month, day, false)
	}
}

func (sch *ScheduleTime) addYMDHM(year *int, month *time.Month, day *int, hour *int, min *int) bool {
	if sch.Minute == nil {
		// minute can be manipulated
		if *min == 59 {
			if !sch.addYMDH(year, month, day, hour) {
				return false
			}
			*min = 0
		} else {
			*min++
		}

		return true
	} else {
		return sch.addYMDH(year, month, day, hour)
	}
}
