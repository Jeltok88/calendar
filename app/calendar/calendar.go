package calendar

import (
	"fmt"

	"github.com/Jeltok88/calendar/app/events"
)

var eventsMap = make(map[string]events.Event)

func AddEvent(title string, date string) (events.Event, error) {

	e, err := events.NewEvent(title, date)
	if err != nil {
		return events.Event{}, err
	}

	eventsMap[e.ID] = e
	return e, nil
}

func DeleteEvent(key string) {
	delete(eventsMap, key)
}

func EditEvent(key string, title string, dateStr string) error {
	e, err := events.NewEvent(title, dateStr)
	if err != nil {
		return err
	}
	eventsMap[key] = e
	return nil
}

func ShowEvents() {
	for _, e := range eventsMap {
		fmt.Println(e.Title, "-", e.StartAt.Format("2006-01-02 15:04"))
	}
}
