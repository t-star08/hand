package thisTime

import (
	"fmt"
	"strings"
)

type ThisClock struct {
	Hour	int
	Minutes	int
	Seconds	int
	Format	string
}

func NewThisClock(hour, minutes, seconds int) *ThisClock {
	c := ThisClock{}

	c.Hour = hour
	c.Minutes = minutes
	c.Seconds = seconds
	c.Format = fmt.Sprintf("%02d:%02d:%02d", c.Hour, c.Minutes, c.Seconds)

	return &c
}

func IsValidThisClock(hour, minutes, seconds int) error {
	err := fmt.Errorf("Uninterpretable Time Expression\n\nTime Expression Must Be\n  HH: [00-23]\n  MM: [00-59]\n  SS: [00-59]")
	if hour < 0 || hour > 23 {
		return err
	}
	if minutes < 0 || minutes > 59 {
		return err
	}
	if seconds < 0 || seconds > 59 {
		return err
	}
	return nil
}

func LessThisClock(c *ThisClock, d *ThisClock) bool {
	return CalcSFromHMS(c.Hour, c.Minutes, c.Seconds) - CalcSFromHMS(d.Hour, d.Minutes, d.Seconds) < 0
}

func ParseClock(value string) (*ThisClock, error) {
	switch strings.Count(value, ":") {
		case 1:
			value += ":00"
		case 2:
			break
		default:
			return &ThisClock{}, fmt.Errorf("Uninterpretable Time Expression\n\nTime Expression Must Be \"HH:MM:SS\" or \"HH:MM\"")
	}

	var h, m, s int
	if _, err := fmt.Sscanf(value, "%d:%d:%d", &h, &m, &s); err != nil {
		return &ThisClock{}, fmt.Errorf("Uninterpretable Time Expression\n\nTime Expression Must Be \"HH:MM:SS\" or \"HH:MM\"")
	}
	if err := IsValidThisClock(h, m, s); err != nil {
		return &ThisClock{}, err
	}
	return NewThisClock(h, m, s), nil
}

func CalcHMSFromS(seconds int) (int, int, int) {
	return seconds / 3600, seconds / 60 % 60, seconds % 60
}

func CalcSFromHMS(hour, minutes, seconds int) int {
	return hour * 3600 + minutes * 60 + seconds
}

func (c *ThisClock) Add(d *ThisDuration) *ThisClock {
	seconds := CalcSFromHMS(c.Hour, c.Minutes, c.Seconds) + CalcSFromHMS(d.Hour, d.Minutes, d.Seconds)
	return NewThisClock(CalcHMSFromS(seconds))
}

func (c *ThisClock) Sub(d *ThisClock) *ThisDuration {
	seconds := CalcSFromHMS(c.Hour, c.Minutes, c.Seconds) - CalcSFromHMS(d.Hour, d.Minutes, d.Seconds)
	return NewThisDuration(CalcHMSFromS(seconds))
}
