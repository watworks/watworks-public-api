package services

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	c := Config()
	require.Equal(t, "http://counter-service:80", c.CounterServiceUrl)
}
