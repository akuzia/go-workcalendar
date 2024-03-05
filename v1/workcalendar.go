package workcalendar

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

var ErrStartDateAfterEnd = errors.New("from start date is after finish date")
var ErrCodeDoesNotExist = errors.New("provided code does not exist")

type HolidayListSchema struct {
	DayOff  []string `json:"dayoff"`
	WorkDay []string `json:"workday"`
}

func loadProviderDates(code string) ([]day, error) {
	f, err := os.Open(fmt.Sprintf("./assets/%s.json", code))
	if errors.Is(err, os.ErrNotExist) {
		return nil, ErrCodeDoesNotExist
	} else if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	s := HolidayListSchema{}
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, err
	}

	days := []day{}
	for _, d := range s.DayOff {
		t, err := time.Parse("20060102", d)
		if err != nil {
			return nil, err
		}

		days = append(days, dayFromTime(t, true))
	}

	for _, d := range s.WorkDay {
		t, err := time.Parse("20060102", d)
		if err != nil {
			return nil, err
		}

		days = append(days, dayFromTime(t, true))
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
