package koiot

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPrevYear(t *testing.T) {
	var sch tScheduleTime
	testTime := time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	year := 2023
	sch.year = &year

	rsp, ok := sch.GetPrevTime(testTime)
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
	var sch tScheduleTime
	year := 2023
	mon := time.March
	sch.year = &year
	sch.month = &mon

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

	sch.year = nil
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
	var sch tScheduleTime
	year := 2023
	mon := time.March
	day := 4
	sch.year = &year
	sch.month = &mon
	sch.day = &day

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

	sch.year = nil
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

	sch.year = &year
	sch.month = nil
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

	sch.year = nil
	sch.month = nil // already is
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
	var sch tScheduleTime
	year := 2023
	mon := time.March
	weekday := time.Saturday
	sch.year = &year
	sch.month = &mon
	sch.weekday = &weekday

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

	sch.year = nil
	testTime = time.Date(2023, 3, 4, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, testTime, rsp)

	testTime = time.Date(2023, 3, 1, 15, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2022, 3, 26, 23, 59, 59, 0, time.UTC), rsp)

	sch.year = &year
	sch.month = nil
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

	sch.year = nil
	sch.month = nil
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
	var sch tScheduleTime
	year := 2023
	mon := time.March
	day := 4
	hour := 15
	sch.year = &year
	sch.month = &mon
	sch.day = &day
	sch.hour = &hour

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

	sch.year = nil
	testTime = time.Date(2023, 3, 4, 14, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2022, 3, 4, 15, 59, 59, 0, time.UTC), rsp)

	sch.year = &year
	sch.month = nil
	testTime = time.Date(2023, 3, 4, 14, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 2, 4, 15, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 1, 4, 14, 30, 10, 0, time.UTC)
	_, ok = sch.GetPrevTime(testTime)
	require.False(t, ok)

	sch.year = nil
	sch.month = nil
	testTime = time.Date(2023, 1, 4, 14, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2022, 12, 4, 15, 59, 59, 0, time.UTC), rsp)

	sch.year = &year
	sch.month = &mon
	sch.day = nil
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

	sch.month = nil
	sch.day = nil
	testTime = time.Date(2023, 3, 1, 14, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 2, 28, 15, 59, 59, 0, time.UTC), rsp)

	sch.year = nil
	sch.month = nil
	day = 31
	sch.day = &day
	testTime = time.Date(2023, 3, 31, 14, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 2, 28, 15, 59, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 1, 14, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 2, 28, 15, 59, 59, 0, time.UTC), rsp)

	sch.day = nil
	weekDay := time.Saturday
	sch.weekday = &weekDay
	testTime = time.Date(2023, 3, 4, 14, 30, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 2, 25, 15, 59, 59, 0, time.UTC), rsp)

}

func TestPrevMinute(t *testing.T) {
	var sch tScheduleTime
	year := 2023
	mon := time.March
	day := 4
	hour := 15
	min := 30
	sch.year = &year
	sch.month = &mon
	sch.day = &day
	sch.hour = &hour
	sch.minute = &min

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

	sch.year = nil
	testTime = time.Date(2023, 3, 4, 15, 29, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2022, 3, 4, 15, 30, 59, 0, time.UTC), rsp)

	sch.year = &year
	sch.hour = nil
	testTime = time.Date(2023, 3, 4, 15, 29, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 14, 30, 59, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 4, 0, 29, 10, 0, time.UTC)
	_, ok = sch.GetPrevTime(testTime)
	require.False(t, ok)

	sch.day = nil
	sch.hour = nil
	testTime = time.Date(2023, 3, 4, 0, 29, 10, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 3, 23, 30, 59, 0, time.UTC), rsp)
}

func TestPrevSecond(t *testing.T) {
	var sch tScheduleTime
	year := 2023
	mon := time.March
	day := 4
	hour := 15
	min := 30
	sec := 10
	sch.year = &year
	sch.month = &mon
	sch.day = &day
	sch.hour = &hour
	sch.minute = &min
	sch.second = &sec

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

	sch.minute = nil
	testTime = time.Date(2023, 3, 4, 15, 30, 9, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 15, 29, 10, 0, time.UTC), rsp)

	testTime = time.Date(2023, 3, 4, 15, 0, 9, 0, time.UTC)
	_, ok = sch.GetPrevTime(testTime)
	require.False(t, ok)

	sch.hour = nil
	sch.minute = nil
	testTime = time.Date(2023, 3, 4, 15, 0, 9, 0, time.UTC)
	rsp, ok = sch.GetPrevTime(testTime)
	require.True(t, ok)
	require.Equal(t, time.Date(2023, 3, 4, 14, 59, 10, 0, time.UTC), rsp)
}