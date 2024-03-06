package workcalendar_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/akuzia/workcalendar"
	"github.com/akuzia/workcalendar/providers"
)

func TestWorkCalendar__SingleDayCheck(t *testing.T) {
	wc, err := workcalendar.NewWorkCalendar("ru")
	require.NoError(t, err)

	assert.False(t, wc.IsWorkday(time.Date(2024, 1, 4, 0, 0, 0, 0, time.UTC)))
	assert.True(t, wc.IsDayOff(time.Date(2024, 1, 4, 0, 0, 0, 0, time.UTC)))

	assert.True(t, wc.IsWorkday(time.Date(2024, 2, 21, 0, 0, 0, 0, time.UTC)))
	assert.False(t, wc.IsDayOff(time.Date(2024, 2, 21, 0, 0, 0, 0, time.UTC)))

	assert.False(t, wc.IsWorkday(time.Date(2024, 3, 16, 0, 0, 0, 0, time.UTC)))
	assert.True(t, wc.IsDayOff(time.Date(2024, 3, 16, 0, 0, 0, 0, time.UTC)))
}

func TestWorkCalendar__WrongCode(t *testing.T) {
	_, err := workcalendar.NewWorkCalendar("oops")
	require.Equal(t, err, providers.ErrCodeDoesNotExist)
}

func TestWorkCalendar__ListDays(t *testing.T) {
	wc, err := workcalendar.NewWorkCalendar("ru")
	require.NoError(t, err)

	holidays, err := wc.ListHolidays(
		time.Date(2023, 12, 30, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
	)
	require.NoError(t, err)
	assert.Equal(t, []time.Time{
		time.Date(2023, 12, 30, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 4, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 6, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 7, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC),
	}, holidays)

	workdays, err := wc.ListWorkdays(
		time.Date(2023, 12, 30, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
	)
	require.NoError(t, err)
	assert.Equal(t, []time.Time{
		time.Date(2024, 1, 9, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
	}, workdays)

}
