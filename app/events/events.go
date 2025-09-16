package events

import (
	"errors"
	"regexp"
	"time"

	"github.com/araddon/dateparse"
	"github.com/google/uuid"
)

type Event struct {
	ID      string
	Title   string
	StartAt time.Time
}

func getNextID() string {
	return uuid.New().String()
}

func NewEvent(title string, dateStr string) (Event, error) {
	t, err := dateparse.ParseAny(dateStr)
	validate := IsValidTitle(title)

	if err != nil {
		return Event{}, errors.New("Неверный формат даты!")
	}

	if !validate {
		return Event{}, errors.New("Некорректное имя задачи!")
	}

	return Event{
		ID:      getNextID(),
		Title:   title,
		StartAt: t,
	}, nil
}

func IsValidTitle(title string) bool {
	pattern := `^[А-Яа-яЁё0-9 ,\.]{3,40}$`
	matched, err := regexp.MatchString(pattern, title)
	if err != nil {
		return false
	}
	return matched
}
