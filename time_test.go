package gotime

import (
    "testing"
    "time"
)

func TestFreeze(t *testing.T) {
    freezeNow := Now()
    Freeze(freezeNow)

    time.Sleep(1 * time.Second)

    now := Now()
    if now.Unix() != freezeNow.Unix() {
        t.Fail()
    }
}
