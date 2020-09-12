package radio

import (
	"sort"
	"time"
)

/* Listener is function which is executed when a eventName is dispatched through a channel */
type Listener func(event Event)

type eventListener struct {
	eventName string
	listener  Listener
	priority  Priority
}

// Adds an listener on any eventName at this or on all parent channels
func (c *Channel) OnEvent(listener Listener, priority Priority) {
	c.listeners = append(c.listeners, &eventListener{
		listener: listener,
		priority: priority,
	})
	sort.Slice(c.listeners, func(i, j int) bool {
		return c.listeners[i].priority > c.listeners[j].priority
	})
}

// Adds an listener on a specific eventName at this or on all parent channels
func (c *Channel) OnEventSpecific(eventName string, listener Listener, priority Priority) {
	c.listeners = append(c.listeners, &eventListener{
		eventName: eventName,
		listener:  listener,
		priority:  priority,
	})
	sort.Slice(c.listeners, func(i, j int) bool {
		return c.listeners[i].priority > c.listeners[j].priority
	})
}

// Does fire an Event on the Channel and all sub-Channels
func (c *Channel) FireEvent(eventName string, payload interface{}) {
	event := Event{
		name:      eventName,
		createdAt: time.Now(),
		payload:   payload,
	}

	c.dispatchEvent(event)
}

func (c *Channel) dispatchEvent(event Event) {
	for _, l := range c.listeners {
		if l.eventName == event.name {
			go l.listener(event)
		} else if l.eventName == "" {
			go l.listener(event)
		}
	}

	for _, child := range c.children {
		go child.dispatchEvent(event)
	}
}
