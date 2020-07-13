package ske

import "reflect"

type Event interface {}

type EventManager struct {
	// store the listeners in a map of the types
	Listeners map[string][]func(event Event)
}

func (e *EventManager) Send(event Event){
	listeners := e.Listeners[reflect.TypeOf(event).String()]
	for _, listener := range listeners{
		listener(event)
	}
}

func (e *EventManager) Listen(tag string, callback func(event Event)){
	e.Listeners[tag] = append(e.Listeners[tag], callback)
}