package radio

import (
	"errors"
	"fmt"
	"time"
)

/* Represent a Channel where you can subscribe to. */
type Channel struct {
	name      string
	listeners []*Listener
}

/* Listener is function which is executed when a event is dispatched through a channel */
type Listener func(event *Event)

/* Event holds the basic data structure which can be passed through the channels */
type Event struct {
	name      string
	createdAt time.Time
	payload   interface{}
}

var channelList []*Channel

func NewChannel(name string) (*Channel, error) {
	if isChannelExisting(name) {
		return nil, errors.New(fmt.Sprintf("The channel %s is already existing, try listen to it instead.", name))
	}

	channel := &Channel{name, []*Listener{}}

	registerChannel(channel)

	return channel, nil
}

func (c *Channel) GetName() string {
	return c.name
}

func registerChannel(newChannel *Channel) {
	channelList = append(channelList, newChannel)
}

func isChannelExisting(channelName string) bool {
	var isExisting bool
	for _, channel := range channelList {
		if channel.name == channelName {
			isExisting = true
			break
		}
	}

	return isExisting
}
