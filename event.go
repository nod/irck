package irck

import (
    "time"
)

type Event struct {
    Origin string
    Body string
    Author string
    ts time.Time
}

const layoutISO = "2006-01-02T15:04:05-0700"


func ParseISOTime(isots string) time.Time {
    t,_ := time.Parse(layoutISO, isots)
    return t
}

func MakeEvent(origin string, body string, author string, ts time.Time) Event {
    return Event{origin, body, author, ts}
}
