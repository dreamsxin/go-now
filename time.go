package now

import "time"

func formatTimeToList(t time.Time) []int {
	hour, min, sec := t.Clock()
	year, month, day := t.Date()
	return []int{t.Nanosecond(), sec, min, hour, day, int(month), year}
}

// Iter is a timerange iterator.
type Iter struct {
	start    time.Time
	end      time.Time
	interval time.Duration
	current  time.Time
	index    int
}

// New creates an Iter.
func NewIter(start time.Time, end time.Time, interval time.Duration) *Iter {
	return &Iter{
		start:    start,
		end:      end,
		interval: interval,
		current:  start,
		index:    0,
	}
}

// Next scans for the next time.
func (i *Iter) Next() bool {
	var next time.Time
	if i.index == 0 {
		next = i.current
	} else {
		next = i.current.Add(i.interval)
	}

	if i.end.Equal(next) || i.end.After(next) {
		i.current = next
		i.index++
		return true
	}
	return false
}

// Current returns the latest unscanned time.
func (i *Iter) Current() time.Time {
	return i.current
}
