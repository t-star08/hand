package thisTime

import (
	"fmt"
)

type ThisDuration struct {
	Hour	int
	Minutes	int
	Seconds	int
	Format	string
}

func NewThisDuration(hour, minutes, seconds int) *ThisDuration {
	d := ThisDuration{}
	d.Hour = hour
	d.Minutes = minutes
	d.Seconds = seconds
	d.Format = fmt.Sprintf("%02dh%02dm%02ds", d.Hour, d.Minutes, d.Seconds)

	return &d
}

func (d *ThisDuration) Add(e *ThisDuration) *ThisDuration {
	seconds := CalcSFromHMS(d.Hour, d.Minutes, d.Seconds) + CalcSFromHMS(e.Hour, e.Minutes, e.Seconds)
	return NewThisDuration(CalcHMSFromS(seconds))
}

func (d *ThisDuration) Sub(e *ThisDuration) *ThisDuration {
	seconds := CalcSFromHMS(d.Hour, d.Minutes, d.Seconds) - CalcSFromHMS(e.Hour, e.Minutes, e.Seconds)
	return NewThisDuration(CalcHMSFromS(seconds))
}
