package thisTime

import (
	"fmt"
	"time"
)

type ThisTime struct {
	Year			int
	Month			int
	Day				int
	Clock			*ThisClock
	ForLogFileName	string
	ForInLog		string
	Format			string
}

func NewThisTime(u time.Time) *ThisTime {
	t := ThisTime{}

	Y, M, D := u.Date()
	h, m, s := u.Clock()

	t.Year = Y
	t.Month = int(M)
	t.Day = D
	t.Clock = NewThisClock(h, m, s)

	t.ForLogFileName = fmt.Sprintf("%d-%02d-%02d", t.Year, t.Month, t.Day)
	t.ForInLog = t.Clock.Format
	t.Format = fmt.Sprintf("%d/%02d/%02d-%02d:%02d:%02d", t.Year, t.Month, t.Day, t.Clock.Hour, t.Clock.Minutes, t.Clock.Seconds)

	return &t
}

func Lock() *ThisTime {
	return NewThisTime(time.Now())
}

func Parse(value string) (*ThisTime, error) {
	// format : 2021/10/18-21:32:38
	if u, err := time.Parse("2006/01/02-15:04:05", value); err != nil {
		return &ThisTime{}, err
	} else {
		return NewThisTime(u), nil
	}
}
