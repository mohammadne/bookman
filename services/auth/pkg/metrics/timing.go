package metrics

import "time"

type Timing struct {
	StartedAt  time.Time
	FinishedAt time.Time
}

func (t *Timing) Start() {
	t.StartedAt = time.Now()
}

func (t *Timing) Finish() {
	t.FinishedAt = time.Now()
}
func (t *Timing) Duration() float64 {
	return time.Since(t.StartedAt).Seconds()
}
