package main

import (
	"fmt"
	"time"

	"github.com/Jeltok88/calendar/app/calendar"
	"github.com/Jeltok88/calendar/app/events"
)

func main() {
	e := events.Event{
		Title:   "Встреча",
		StartAt: time.Now(),
	}
	calendar.AddEvent("event1", e)
	fmt.Println("Календарь обновлён")
}
