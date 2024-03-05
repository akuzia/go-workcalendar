package workcalendar

import "time"

type day struct {
	time.Time

	hodiday bool
}

func (d day) IsHoliday() bool {
	return d.hodiday
}

func isWeekend(t time.Time) bool {
	return t.Weekday() == time.Sunday || t.Weekday() == time.Saturday
}

func dayFromTime(t time.Time, holiday bool) day {
	return day{
		Time:    t,
		hodiday: holiday,
	}
}

type dayListKey struct {
	year  int
	month time.Month
	day   int
}

func listKey(t time.Time) dayListKey {
	return dayListKey{
		year:  t.Year(),
		month: t.Month(),
		day:   t.Day(),
	}
}

func rangeDate(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, start.Location())
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, end.Location())

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}
