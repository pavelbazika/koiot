package koiot

import "time"

type (
	tScheduleTime struct {
		year    *int
		month   *time.Month
		day     *int
		weekday *time.Weekday
		hour    *int
		minute  *int
		second  *int
	}

	tScheduleItem struct {
		time tScheduleTime
		data any
	}

	tSchedule []tScheduleItem
)

func (schedule tSchedule) getCurrentItemData(t time.Time) any {
	// find the nearest record
	winner := -1
	var winnerTime time.Time
	for i, sch := range schedule {
		trigTime, found := sch.time.GetPrevTime(t)
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
		return schedule[winner].data
	} else {
		return nil
	}
}

// get time when the item triggered last before (or at) t
func (sch *tScheduleTime) GetPrevTime(t time.Time) (time.Time, bool) {
	if (sch.weekday != nil) && (sch.day != nil) {
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
	if sch.year != nil {
		yr = *sch.year
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
	case sch.month != nil:
		mon = *sch.month

	case prevTime:
		mon = 12

	default:
		mon = rqMon
	}
	if !prevTime && (mon > rqMon) {
		// month specified by schedule, after rq
		if sch.year == nil {
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
	case sch.day != nil:
		day = *sch.day

	case prevTime:
		day = daysIn(yr, mon)

	default:
		day = rqDay
	}
	if sch.weekday != nil {
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
	case sch.hour != nil:
		hr = *sch.hour

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
	case sch.minute != nil:
		min = *sch.minute

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
	case sch.second != nil:
		sec = *sch.second

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
func (sch *tScheduleTime) GetNextTime(t time.Time) (time.Time, bool) {
	if (sch.weekday != nil) && (sch.day != nil) {
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
	if sch.year != nil {
		yr = *sch.year
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
	case sch.month != nil:
		mon = *sch.month

	case prevTime:
		mon = 12

	default:
		mon = rqMon
	}
	if !prevTime && (mon > rqMon) {
		// month specified by schedule, after rq
		if sch.year == nil {
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
	case sch.day != nil:
		day = *sch.day

	case prevTime:
		day = daysIn(yr, mon)

	default:
		day = rqDay
	}
	if sch.weekday != nil {
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
	case sch.hour != nil:
		hr = *sch.hour

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
	case sch.minute != nil:
		min = *sch.minute

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
	case sch.second != nil:
		sec = *sch.second

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

func daysIn(year int, m time.Month) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

// subtract month to previous ocurrence
func (sch *tScheduleTime) subYM(year *int, month *time.Month) bool {
	if sch.month == nil {
		// month can be manipulated
		if *month == time.January {
			if sch.year == nil {
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
		if sch.year == nil {
			*year--
		} else {
			return false
		}
	}

	return true
}

func (sch *tScheduleTime) subYMD(year *int, month *time.Month, day *int, weekDayFind bool) bool {
	if sch.day == nil {
		if sch.weekday == nil {
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
			if dayWd == *sch.weekday {
				// we are on requested weekday
				if !weekDayFind {
					diff = 7
				}
			} else {
				diff = int(dayWd) - int(*sch.weekday)
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
				if dayWd != *sch.weekday {
					diff = int(dayWd) - int(*sch.weekday)
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

func (sch *tScheduleTime) subYMDH(year *int, month *time.Month, day *int, hour *int) bool {
	if sch.hour == nil {
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

func (sch *tScheduleTime) subYMDHM(year *int, month *time.Month, day *int, hour *int, min *int) bool {
	if sch.minute == nil {
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
