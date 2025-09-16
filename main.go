package main

import (
	"fmt"

	"github.com/Jeltok88/calendar/app/calendar"
)

func main() {

	_, err1 := calendar.AddEvent("Митинг", "2025-09-01 09:30")
	if err1 != nil {
		fmt.Println("Ошибка", err1)
		return
	}
	event2, err2 := calendar.AddEvent("Пробежка", "2025-08-01 12:30")
	if err2 != nil {
		fmt.Println("Ошибка", err2)
		return
	}
	event3, err3 := calendar.AddEvent("Откликнуться на вакансии", "2025-08-02 7:30")
	if err3 != nil {
		fmt.Println("Ошибка", err3)
		return
	}
	calendar.ShowEvents()

	calendar.DeleteEvent(event2.ID)
	err := calendar.EditEvent(event3.ID, "Театр", "2025-08-02 10:00")
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	calendar.ShowEvents()

	fmt.Scanln()
}
