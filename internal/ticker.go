package internal

import (
	"time"
)

type Ticker struct {
	next time.Time
	dur  time.Duration
	C    <-chan struct{}
}

func (t *Ticker) tick(ch chan struct{}) {
	t.next = time.Now().Add(t.dur)
	ch <- struct{}{}
}

func (t *Ticker) Stop() {
}

func NewTicker(dur time.Duration) *Ticker {
	ch := make(chan struct{})
	ticker := &Ticker{
		next: time.Now().Add(dur),
		dur:  dur,
		C:    ch}
	go func() {
		defer func() { close(ch) }()
		for {
			select {
			// case <-ctx.Done():
			// 	return
			default:
				now := time.Now()
				if !now.After(ticker.next) {
					time.Sleep(ticker.next.Sub(now))
				}
				ticker.tick(ch)
			}
		}
	}()
	return ticker
}
