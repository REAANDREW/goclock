package goclock

import (
	"github.com/franela/goblin"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Clock(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Should time warp in units of:", func() {
		var clock *FakeClock

		g.BeforeEach(func() {
			clock = NewFakeClock()
		})

		g.It("Milliseconds", func() {
			start := clock.Now()
			clock.Tick("500ms")
			end := clock.Now()
			assert.Equal(t, 500*time.Millisecond, DurationDiff(start, end), "duration difference not expected")
		})

		g.It("Seconds", func() {
			start := clock.Now()
			clock.Tick("500s")
			end := clock.Now()
			assert.Equal(t, 500*time.Second, DurationDiff(start, end), "duration difference not expected")
		})

		g.It("Minutes", func() {
			start := clock.Now()
			clock.Tick("500m")
			end := clock.Now()
			assert.Equal(t, 500*time.Minute, DurationDiff(start, end), "duration difference not expected")
		})
	})
}
