package main

import (
	"fmt"
	"time"

	"github.com/Jeltok88/app/calendar"
	"github.com/Jeltok88/app/events"
)

func main() {
	e := events.Event{
		Title:   "Встреча",
		StartAt: time.Now(),
	}
	calendar.AddEvent("event1", e)
	fmt.Println("Календарь обновлён")
}
