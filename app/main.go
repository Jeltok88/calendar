package main

import (
    "fmt"
    "github.com/Jeltok88/calendar/app/calendar"
    "github.com/Jeltok88/calendar/app/events"
    "time"
)

func main() {
    e := events.Event{
        Title:   "Встреча",
        StartAt: time.Now(),
    }
    calendar.AddEvent("event1", e)
    fmt.Println("Календарь обновлён")
}