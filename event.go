package ske

import "reflect"

type Event interface {}

type EventManager struct {
	// store the listeners in a map of the types
	Listeners map[string][]func(event Event)
}

// send an event to the event bus
func (e *EventManager) Send(event Event){
	listeners := e.Listeners[reflect.TypeOf(event).String()]
	for _, listener := range listeners{
		listener(event)
	}
}

// listen to events of a particular type
func (e *EventManager) Listen(tag string, callback func(event Event)){
	e.Listeners[tag] = append(e.Listeners[tag], callback)
}