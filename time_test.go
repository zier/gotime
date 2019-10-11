package gotime

import (
    "testing"
    "time"

    "github.com/stretchr/testify/require"
)

func TestFreeze(t *testing.T) {
    freezeNow := Now()
    Freeze(freezeNow)

    time.Sleep(1 * time.Second)

    now := Now()
    require.Equal(t, freezeNow, now)
}
