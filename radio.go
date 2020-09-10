package radio

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

type Radio struct {
	channelList []*Channel
}

/* Listener is function which is executed when a eventName is dispatched through a channel */
type Listener func(event Event)

/* Event holds the basic data structure which can be passed through the channels */
type Event struct {
	name      string
	createdAt time.Time
	payload   interface{}
}

func NewRadio() (*Radio, error) {
	return new(Radio), nil
}

// Create a new Channel to work with it
func (r *Radio) NewChannel(path string) (*Channel, error) {
	if r.FindChannel(path) != nil {
		return nil, errors.New(fmt.Sprintf("The channel with the path %s is already existing, try find to it instead.", path))
	}

	parentChannelPath := getParentChannelPath(path)
	var parentChannel = r.findParentChannel(path)
	if parentChannel == nil && parentChannelPath != "" {
		parentChannel, _ = r.NewChannel(parentChannelPath)
	}

	channel := &Channel{
		parent:     parentChannel,
		path:       path,
		dispatcher: func(event Event) {},
	}

	r.registerChannel(channel)

	if parentChannel != nil {
		parentChannel.addChildren(channel)
	}

	return channel, nil
}

// FindChannel does helps you finding the right Channel by it's path
func (r *Radio) FindChannel(path string) *Channel {
	var foundChannel *Channel
	for _, channel := range r.channelList {
		if channel.path == path {
			foundChannel = channel
			break
		}
	}

	return foundChannel
}

// PRIVATE FUNC's

func getParentChannelPath(channelPath string) string {
	splicedPath := strings.Split(channelPath, DefaultPathSeparator)
	if len(splicedPath) <= 1 {
		return ""
	}

	return strings.Join(splicedPath[:(len(splicedPath)-1)], DefaultPathSeparator)
}

func (r *Radio) registerChannel(newChannel *Channel) {
	r.channelList = append(r.channelList, newChannel)
}
func (r *Radio) findParentChannel(channelPath string) *Channel {
	parentPath := getParentChannelPath(channelPath)
	return r.FindChannel(parentPath)
}

func (c *Channel) addChildren(newChildren *Channel) {
	if c.path == newChildren.path {
		// Internally, that the same channel can never be added as a children to itself
		log.Fatalln(fmt.Sprintf("You can't add the same channel to your own channel"))
	}

	c.children = append(c.children, newChildren)
}
