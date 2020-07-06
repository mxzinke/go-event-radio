package radio

import (
	"sort"
	"time"
)

type eventListener struct {
	eventName string
	listener  Listener
	priority  Priority
}

type eventDispatcher func(event Event)

// Adds an listener on any eventName at this or on all parent channels
func (c *Channel) OnEvent(listener Listener, priority Priority) {
	c.listeners = append(c.listeners, &eventListener{
		listener: listener,
		priority: priority,
	})
}

// Adds an listener on a specific eventName at this or on all parent channels
func (c *Channel) OnEventSpecific(eventName string, listener Listener, priority Priority) {
	c.listeners = append(c.listeners, &eventListener{
		eventName: eventName,
		listener:  listener,
		priority:  priority,
	})
	c.dispatcher = c.buildEventDispatcher()
}

func (c *Channel) FireEvent(eventName string, payload interface{}) {
	event := Event{
		name:      eventName,
		createdAt: time.Now(),
		payload:   payload,
	}

	c.dispatcher(event)
}

func (c *Channel) buildEventDispatcher() eventDispatcher {
	listenersList := append(c.listeners, c.getChildrenListeners()...)
	sort.Slice(listenersList, func(i, j int) bool {
		return listenersList[i].priority > listenersList[j].priority
	})

	return func(event Event) {
		for _, l := range listenersList {
			if l.eventName == event.name {
				go l.listener(event)
			} else if l.eventName == "" {
				go l.listener(event)
			}
		}
	}
}

func (c *Channel) getChildrenListeners() []*eventListener {
	var allListeners []*eventListener
	for _, channel := range c.children {
		if len(channel.children) > 0 {
			allListeners = append(allListeners, channel.getChildrenListeners()...)
			allListeners = append(allListeners, channel.listeners...)
		} else {
			allListeners = append(allListeners, channel.listeners...)
		}
	}

	return allListeners
}
