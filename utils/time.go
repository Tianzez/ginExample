package utils

import "time"

var (
	DateFormat = "2006-01-02"
	TimeFormat = "2006-01-02 00:00:00"
)

func RangeDate(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}

func PreDate() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

func PreDateString() string {
	return PreDate().Format(DateFormat)
}
