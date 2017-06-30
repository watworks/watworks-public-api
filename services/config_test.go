package services

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	c := Config()
	require.Equal(t, "http://counter-service:8006", c.CounterServiceUrl)
}
