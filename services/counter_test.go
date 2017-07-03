package services

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCounterService(t *testing.T) {
	r := require.New(t)
	s := Counter()

	e := func() string {
		if err := s.Err(); err != nil {
			return err.Error()
		}

		return ""
	}

	// test init, reset, incr/decr behavior
	r.True(s.Init(), e())
	r.True(s.Reset(), e())
	r.Equal(float64(0), s.Get(), e())
	r.True(s.Incr(), e())
	r.True(s.Incr(), e())
	r.Equal(float64(2), s.Get(), e())
	r.True(s.Decr(), e())
	r.Equal(float64(1), s.Get(), e())
	r.True(s.Reset(), e())
	r.Equal(float64(0), s.Get(), e())
}
