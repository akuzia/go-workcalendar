package providers

import "time"

func newRUProvider() Provider {
	return Provider{
		DaysOff: []time.Time{
			time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
			time.Date(2024, 1, 2, 0, 0, 0, 0, time.Local),
			time.Date(2024, 1, 3, 0, 0, 0, 0, time.Local),
			time.Date(2024, 1, 4, 0, 0, 0, 0, time.Local),
			time.Date(2024, 1, 5, 0, 0, 0, 0, time.Local),
			time.Date(2024, 1, 8, 0, 0, 0, 0, time.Local),
			time.Date(2024, 2, 23, 0, 0, 0, 0, time.Local),
			time.Date(2024, 3, 8, 0, 0, 0, 0, time.Local),
			time.Date(2024, 4, 29, 0, 0, 0, 0, time.Local),
			time.Date(2024, 4, 30, 0, 0, 0, 0, time.Local),
			time.Date(2024, 5, 1, 0, 0, 0, 0, time.Local),
			time.Date(2024, 5, 9, 0, 0, 0, 0, time.Local),
			time.Date(2024, 5, 10, 0, 0, 0, 0, time.Local),
			time.Date(2024, 6, 12, 0, 0, 0, 0, time.Local),
			time.Date(2024, 11, 4, 0, 0, 0, 0, time.Local),
			time.Date(2024, 12, 30, 0, 0, 0, 0, time.Local),
			time.Date(2024, 12, 31, 0, 0, 0, 0, time.Local),
		},
		WorkDays: []time.Time{
			time.Date(2024, 04, 27, 0, 0, 0, 0, time.Local),
			time.Date(2024, 11, 02, 0, 0, 0, 0, time.Local),
			time.Date(2024, 12, 28, 0, 0, 0, 0, time.Local),
		},
	}
}
