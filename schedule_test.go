package koiot

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPrevYear(t *testing.T) {
	var sch ScheduleTime
	testTime := time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	year := 2023
	sch.Year = &year

	rsp, ok := sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 3, 4, 13, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2024, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2022, 3, 4, 15, 30, 10, 0, time.UTC)
	_, ok = sch.GetPrevTime(testTime)
	require.False(t, ok)
}

func TestPrevMonth(t *testing.T) {
	var sch ScheduleTime
	year := 2023
	mon := time.March
	sch.Year = &year
	sch.Month = &mon

	testTime := time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok := sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 4, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 31, 23, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2024, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 31, 23, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 2, 4, 15, 30, 10, 0, time.UTC)
	_, ok = sch.GetPrevTime(testTime)
	require.False(t, ok)

	sch.Year = nil
	testTime = time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 2, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2022, 3, 31, 23, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 4, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 31, 23, 59, 59, 0, time.UTC), rsp)
}

func TestPrevDay(t *testing.T) {
	var sch ScheduleTime
	year := 2023
	mon := time.March
	day := 4
	sch.Year = &year
	sch.Month = &mon
	sch.Day = &day

	testTime := time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok := sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 4, 5, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 23, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 2, 2, 15, 30, 10, 0, time.UTC)
	_, ok = sch.GetPrevTime(testTime)
	require.False(t, ok)

	testTime = time.Date(2023, 3, 2, 15, 30, 10, 0, time.UTC)
	_, ok = sch.GetPrevTime(testTime)
	require.False(t, ok)

	sch.Year = nil
	testTime = time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 1, 1, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2022, 3, 4, 23, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 4, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 23, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 2, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2022, 3, 4, 23, 59, 59, 0, time.UTC), rsp)

	sch.Year = &year
	sch.Month = nil
	testTime = time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 5, 1, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 4, 4, 23, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 5, 7, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 5, 4, 23, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2022, 1, 1, 15, 30, 10, 0, time.UTC)
	_, ok = sch.GetPrevTime(testTime)
	require.False(t, ok)

	testTime = time.Date(2023, 1, 2, 15, 30, 10, 0, time.UTC)
	_, ok = sch.GetPrevTime(testTime)
	require.False(t, ok)

	sch.Year = nil
	sch.Month = nil // already is
	testTime = time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 5, 1, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 4, 4, 23, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 1, 1, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2022, 12, 4, 23, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 5, 7, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 5, 4, 23, 59, 59, 0, time.UTC), rsp)
}

func TestPrevWeekDay(t *testing.T) {
	var sch ScheduleTime
	year := 2023
	mon := time.March
	weekday := time.Saturday
	sch.Year = &year
	sch.Month = &mon
	sch.WeekDay = &weekday

	testTime := time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok := sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 3, 5, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 23, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 6, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 23, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 7, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 23, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 8, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 23, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 9, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 23, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 10, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 23, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 11, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 11, 15, 30, 10, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 12, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 11, 23, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 3, 15, 30, 10, 0, time.UTC)
	_, ok = sch.GetPrevTime(testTime)
	require.False(t, ok)

	sch.Year = nil
	testTime = time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 3, 1, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2022, 3, 26, 23, 59, 59, 0, time.UTC), rsp)

	sch.Year = &year
	sch.Month = nil
	testTime = time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 3, 1, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 2, 25, 23, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 1, 7, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 1, 7, 15, 30, 10, 0, time.UTC), rsp)

	testTime = time.Date(2023, 1, 6, 15, 30, 10, 0, time.UTC)
	_, ok = sch.GetPrevTime(testTime)
	require.False(t, ok)

	sch.Year = nil
	sch.Month = nil
	testTime = time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 1, 6, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2022, 12, 31, 23, 59, 59, 0, time.UTC), rsp)
}

func TestPrevHour(t *testing.T) {
	var sch ScheduleTime
	year := 2023
	mon := time.March
	day := 4
	hour := 15
	sch.Year = &year
	sch.Month = &mon
	sch.Day = &day
	sch.Hour = &hour

	testTime := time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok := sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 4, 5, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 15, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 4, 14, 30, 10, 0, time.UTC)
	_, ok = sch.GetPrevTime(testTime)
	require.False(t, ok)

	sch.Year = nil
	testTime = time.Date(2023, 3, 4, 14, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2022, 3, 4, 15, 59, 59, 0, time.UTC), rsp)

	sch.Year = &year
	sch.Month = nil
	testTime = time.Date(2023, 3, 4, 14, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 2, 4, 15, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 1, 4, 14, 30, 10, 0, time.UTC)
	_, ok = sch.GetPrevTime(testTime)
	require.False(t, ok)

	sch.Year = nil
	sch.Month = nil
	testTime = time.Date(2023, 1, 4, 14, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2022, 12, 4, 15, 59, 59, 0, time.UTC), rsp)

	sch.Year = &year
	sch.Month = &mon
	sch.Day = nil
	testTime = time.Date(2023, 3, 4, 14, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 3, 15, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 4, 16, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 15, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 1, 14, 30, 10, 0, time.UTC)
	_, ok = sch.GetPrevTime(testTime)
	require.False(t, ok)

	sch.Month = nil
	sch.Day = nil
	testTime = time.Date(2023, 3, 1, 14, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 2, 28, 15, 59, 59, 0, time.UTC), rsp)

	sch.Year = nil
	sch.Month = nil
	day = 31
	sch.Day = &day
	testTime = time.Date(2023, 3, 31, 14, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 2, 28, 15, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 1, 14, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 2, 28, 15, 59, 59, 0, time.UTC), rsp)

	sch.Day = nil
	weekDay := time.Saturday
	sch.WeekDay = &weekDay
	testTime = time.Date(2023, 3, 4, 14, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 2, 25, 15, 59, 59, 0, time.UTC), rsp)

}

func TestPrevMinute(t *testing.T) {
	var sch ScheduleTime
	year := 2023
	mon := time.March
	day := 4
	hour := 15
	min := 30
	sch.Year = &year
	sch.Month = &mon
	sch.Day = &day
	sch.Hour = &hour
	sch.Minute = &min

	testTime := time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok := sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 4, 5, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 15, 30, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 4, 15, 29, 10, 0, time.UTC)
	_, ok = sch.GetPrevTime(testTime)
	require.False(t, ok)

	sch.Year = nil
	testTime = time.Date(2023, 3, 4, 15, 29, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2022, 3, 4, 15, 30, 59, 0, time.UTC), rsp)

	sch.Year = &year
	sch.Hour = nil
	testTime = time.Date(2023, 3, 4, 15, 29, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 14, 30, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 4, 0, 29, 10, 0, time.UTC)
	_, ok = sch.GetPrevTime(testTime)
	require.False(t, ok)

	sch.Day = nil
	sch.Hour = nil
	testTime = time.Date(2023, 3, 4, 0, 29, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 3, 23, 30, 59, 0, time.UTC), rsp)
}

func TestPrevSecond(t *testing.T) {
	var sch ScheduleTime
	year := 2023
	mon := time.March
	day := 4
	hour := 15
	min := 30
	sec := 10
	sch.Year = &year
	sch.Month = &mon
	sch.Day = &day
	sch.Hour = &hour
	sch.Minute = &min
	sch.Second = &sec

	testTime := time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok := sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 3, 4, 15, 30, 11, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 4, 15, 30, 9, 0, time.UTC)
	_, ok = sch.GetPrevTime(testTime)
	require.False(t, ok)

	sch.Minute = nil
	testTime = time.Date(2023, 3, 4, 15, 30, 9, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 15, 29, 10, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 4, 15, 0, 9, 0, time.UTC)
	_, ok = sch.GetPrevTime(testTime)
	require.False(t, ok)

	sch.Hour = nil
	sch.Minute = nil
	testTime = time.Date(2023, 3, 4, 15, 0, 9, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 14, 59, 10, 0, time.UTC), rsp)
}

func TestNextYear(t *testing.T) {
	var sch ScheduleTime
	testTime := time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	year := 2023
	sch.Year = &year

	rsp, ok := sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 3, 4, 18, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2022, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2024, 3, 4, 15, 30, 10, 0, time.UTC)
	_, ok = sch.GetNextTime(testTime)
	require.False(t, ok)
}

func TestNextMonth(t *testing.T) {
	var sch ScheduleTime
	year := 2023
	mon := time.March
	sch.Year = &year
	sch.Month = &mon

	testTime := time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok := sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 2, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2022, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 4, 4, 15, 30, 10, 0, time.UTC)
	_, ok = sch.GetNextTime(testTime)
	require.False(t, ok)

	sch.Year = nil
	testTime = time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 4, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 2, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC), rsp)
}

func TestNextDay(t *testing.T) {
	var sch ScheduleTime
	year := 2023
	mon := time.March
	day := 4
	sch.Year = &year
	sch.Month = &mon
	sch.Day = &day

	testTime := time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok := sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 2, 3, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 0, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 4, 6, 15, 30, 10, 0, time.UTC)
	_, ok = sch.GetNextTime(testTime)
	require.False(t, ok)

	testTime = time.Date(2023, 3, 6, 15, 30, 10, 0, time.UTC)
	_, ok = sch.GetNextTime(testTime)
	require.False(t, ok)

	sch.Year = nil
	testTime = time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 12, 31, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2024, 3, 4, 0, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 2, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 0, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 6, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2024, 3, 4, 0, 0, 0, 0, time.UTC), rsp)

	sch.Year = &year
	sch.Month = nil
	testTime = time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 1, 7, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 2, 4, 0, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 5, 2, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 5, 4, 0, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2024, 1, 1, 15, 30, 10, 0, time.UTC)
	_, ok = sch.GetNextTime(testTime)
	require.False(t, ok)

	testTime = time.Date(2023, 12, 6, 15, 30, 10, 0, time.UTC)
	_, ok = sch.GetNextTime(testTime)
	require.False(t, ok)

	sch.Year = nil
	sch.Month = nil // already is
	testTime = time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 2, 28, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 0, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 12, 31, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2024, 1, 4, 0, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 5, 1, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 5, 4, 0, 0, 0, 0, time.UTC), rsp)
}

func TestNextWeekDay(t *testing.T) {
	var sch ScheduleTime
	year := 2023
	mon := time.March
	weekday := time.Saturday
	sch.Year = &year
	sch.Month = &mon
	sch.WeekDay = &weekday

	testTime := time.Date(2023, 3, 11, 15, 30, 10, 0, time.UTC)
	rsp, ok := sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 3, 10, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 11, 0, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 9, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 11, 0, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 8, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 11, 0, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 7, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 11, 0, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 6, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 11, 0, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 5, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 11, 0, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 3, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 0, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 31, 15, 30, 10, 0, time.UTC)
	_, ok = sch.GetNextTime(testTime)
	require.False(t, ok)

	sch.Year = nil
	testTime = time.Date(2023, 3, 11, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 3, 31, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2024, 3, 2, 0, 0, 0, 0, time.UTC), rsp)

	sch.Year = &year
	sch.Month = nil
	testTime = time.Date(2023, 3, 11, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 3, 31, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 12, 30, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 12, 30, 15, 30, 10, 0, time.UTC), rsp)

	testTime = time.Date(2023, 12, 31, 15, 30, 10, 0, time.UTC)
	_, ok = sch.GetNextTime(testTime)
	require.False(t, ok)

	sch.Year = nil
	sch.Month = nil
	testTime = time.Date(2023, 3, 11, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 12, 31, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2024, 1, 6, 0, 0, 0, 0, time.UTC), rsp)
}

func TestNextHour(t *testing.T) {
	var sch ScheduleTime
	year := 2023
	mon := time.March
	day := 4
	hour := 15
	sch.Year = &year
	sch.Month = &mon
	sch.Day = &day
	sch.Hour = &hour

	testTime := time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok := sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 2, 3, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 15, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 4, 16, 30, 10, 0, time.UTC)
	_, ok = sch.GetNextTime(testTime)
	require.False(t, ok)

	sch.Year = nil
	testTime = time.Date(2023, 3, 4, 16, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2024, 3, 4, 15, 0, 0, 0, time.UTC), rsp)

	sch.Year = &year
	sch.Month = nil
	testTime = time.Date(2023, 3, 4, 16, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 4, 4, 15, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 12, 4, 16, 30, 10, 0, time.UTC)
	_, ok = sch.GetNextTime(testTime)
	require.False(t, ok)

	sch.Year = nil
	sch.Month = nil
	testTime = time.Date(2023, 12, 4, 16, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2024, 1, 4, 15, 0, 0, 0, time.UTC), rsp)

	sch.Year = &year
	sch.Month = &mon
	sch.Day = nil
	testTime = time.Date(2023, 3, 4, 16, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 5, 15, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 4, 14, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 15, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 31, 16, 30, 10, 0, time.UTC)
	_, ok = sch.GetNextTime(testTime)
	require.False(t, ok)

	sch.Month = nil
	sch.Day = nil
	testTime = time.Date(2023, 3, 31, 16, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 4, 1, 15, 0, 0, 0, time.UTC), rsp)

	sch.Year = nil
	sch.Month = nil
	day = 31
	sch.Day = &day
	testTime = time.Date(2023, 3, 1, 16, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 31, 15, 0, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 31, 16, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 4, 30, 15, 0, 0, 0, time.UTC), rsp)

	sch.Day = nil
	weekDay := time.Saturday
	sch.WeekDay = &weekDay
	testTime = time.Date(2023, 3, 4, 16, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 11, 15, 0, 0, 0, time.UTC), rsp)
}

func TestNextMinute(t *testing.T) {
	var sch ScheduleTime
	year := 2023
	mon := time.March
	day := 4
	hour := 15
	min := 30
	sch.Year = &year
	sch.Month = &mon
	sch.Day = &day
	sch.Hour = &hour
	sch.Minute = &min

	testTime := time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok := sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 2, 3, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 15, 30, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 4, 15, 31, 10, 0, time.UTC)
	_, ok = sch.GetNextTime(testTime)
	require.False(t, ok)

	sch.Year = nil
	testTime = time.Date(2023, 3, 4, 15, 31, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2024, 3, 4, 15, 30, 0, 0, time.UTC), rsp)

	sch.Year = &year
	sch.Hour = nil
	testTime = time.Date(2023, 3, 4, 15, 31, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 16, 30, 0, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 4, 23, 31, 10, 0, time.UTC)
	_, ok = sch.GetNextTime(testTime)
	require.False(t, ok)

	sch.Day = nil
	sch.Hour = nil
	testTime = time.Date(2023, 3, 4, 23, 31, 10, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 5, 0, 30, 0, 0, time.UTC), rsp)
}

func TestNextSecond(t *testing.T) {
	var sch ScheduleTime
	year := 2023
	mon := time.March
	day := 4
	hour := 15
	min := 30
	sec := 10
	sch.Year = &year
	sch.Month = &mon
	sch.Day = &day
	sch.Hour = &hour
	sch.Minute = &min
	sch.Second = &sec

	testTime := time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok := sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 3, 4, 15, 30, 9, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 4, 15, 30, 11, 0, time.UTC)
	_, ok = sch.GetNextTime(testTime)
	require.False(t, ok)

	sch.Minute = nil
	testTime = time.Date(2023, 3, 4, 15, 30, 11, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 15, 31, 10, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 4, 15, 59, 11, 0, time.UTC)
	_, ok = sch.GetNextTime(testTime)
	require.False(t, ok)

	sch.Hour = nil
	sch.Minute = nil
	testTime = time.Date(2023, 3, 4, 15, 59, 11, 0, time.UTC)
	rsp, ok = sch.GetNextTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 16, 0, 10, 0, time.UTC), rsp)
}
