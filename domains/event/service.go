package event

import "log"

type EventService struct {
}

func NewEventService() *EventService {
	instance := &EventService{}
	return instance
}
func (e EventService) GetEventHandler() Event {
	log.Println("event log")
	return e.GetEvent()
}

func (e EventService) GetEvent() Event {
	return Event{}
}
