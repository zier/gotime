package gotime

import "time"

var isFreeze bool
var timeFreeze time.Time

func Freeze(t time.Time) {
    timeFreeze = t
    isFreeze = true
}

func Now() time.Time {
    if isFreeze {
        return timeFreeze
    }

    return time.Now()
}

func NowUnix() int64 {
    if isFreeze {
        return timeFreeze.Unix()
    }

    return time.Now().Unix()
}
