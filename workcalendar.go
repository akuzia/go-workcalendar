package workcalendar

import (
	"errors"
	"time"

	"github.com/akuzia/workcalendar/providers"
)

var ErrStartDateAfterEnd = errors.New("from start date is after finish date")

type HolidayListSchema struct {
	DayOff  []string `json:"dayoff"`
	WorkDay []string `json:"workday"`
}

func loadProviderDates(code string) ([]day, error) {
	days := []day{}
	provider, err := providers.GetProvider(code)
	if err != nil {
		return days, err
	}

	for _, t := range provider.DaysOff {
		days = append(days, dayFromTime(t, true))
	}

	for _, t := range provider.WorkDays {
		days = append(days, dayFromTime(t, false))
	}

	return days, nil
}

type WorkCalendar struct {
	cal map[dayListKey]day
}

func NewWorkCalendar(code string) (*WorkCalendar, error) {
	days, err := loadProviderDates(code)
	if err != nil {
		return nil, err
	}

	cal := make(map[dayListKey]day)
	for _, d := range days {
		cal[listKey(d.Time)] = d
	}

	return &WorkCalendar{
		cal: cal,
	}, nil
}

func (c WorkCalendar) IsWorkday(t time.Time) bool {
	return !c.IsDayOff(t)
}

func (c WorkCalendar) IsDayOff(t time.Time) bool {
	d, ok := c.cal[listKey(t)]
	if ok {
		return d.IsHoliday()
	}

	return isWeekend(t)
}

func (c WorkCalendar) ListWorkdays(
	from time.Time,
	to time.Time,
) ([]time.Time, error) {
	if from.After(to) {
		return nil, ErrStartDateAfterEnd
	}

	days := []time.Time{}

	for d := rangeDate(from, to); ; {
		t := d()

		if t.IsZero() {
			break
		}

		if c.IsWorkday(t) {
			days = append(days, t)
		}
	}

	return days, nil
}

func (c WorkCalendar) ListHolidays(
	from time.Time,
	to time.Time,
) ([]time.Time, error) {
	if from.After(to) {
		return nil, ErrStartDateAfterEnd
	}

	days := []time.Time{}
	for d := rangeDate(from, to); ; {
		t := d()

		if t.IsZero() {
			break
		}

		if c.IsDayOff(t) {
			days = append(days, t)
		}
	}

	return days, nil
}
