package events

import (
	"context"
	"errors"
	"sync"
)

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
	lock     sync.RWMutex
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	ed.lock.Lock()
	defer ed.lock.Unlock()
	exist := ed.Has(eventName, handler)
	if exist {
		return errors.New("handler already exists")
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil

}

func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := ed.handlers[eventName]; ok {
		for _, value := range ed.handlers[eventName] {
			if value == handler {
				return true

			}
		}
	}
	return false
}

func (ed *EventDispatcher) Clear() error {
	ed.lock.Lock()
	defer ed.lock.Unlock()
	ed.handlers = make(map[string][]EventHandlerInterface)
	return nil
}

func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	ed.lock.Lock()
	defer ed.lock.Unlock()
	if _, ok := ed.handlers[eventName]; ok {
		for index, value := range ed.handlers[eventName] {
			if value == handler {
				ed.handlers[eventName] = append(ed.handlers[eventName][:index], ed.handlers[eventName][index+1:]...)
				break
			}
		}
	}
	return nil
}

func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	ed.lock.RLock()
	defer ed.lock.RUnlock()
	wg := &sync.WaitGroup{}
	if handlers, ok := ed.handlers[event.GetName()]; ok {
		for _, handle := range handlers {
			wg.Add(1)
			go handle.Handle(context.Background(), event, wg)
		}
		wg.Wait()
	}
	return nil
}
