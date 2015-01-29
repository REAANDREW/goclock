package goclock

import (
	"fmt"
	"time"
)

type Clock interface {
	Now() time.Time
}

type RealClock struct{}

func (instance *RealClock) Now() (returnTime time.Time) {
	return time.Now()
}

func NewRealClock() (clock *RealClock) {
	return &RealClock{}
}

type FakeClock struct {
	Current time.Time
}

func NewFakeClock() (clock *FakeClock) {
	return &FakeClock{time.Now()}
}

func (instance *FakeClock) Now() (returnTime time.Time) {
	return instance.Current
}

func (instance *FakeClock) Tick(value string) {
	parsedDuration, err := time.ParseDuration(value)
	if err != nil {
		fmt.Errorf("error encountered %v", err)
	}
	instance.Current = instance.Current.Add(parsedDuration)
}

func DurationDiff(start time.Time, end time.Time) (duration time.Duration) {
	return end.Sub(start)
}
