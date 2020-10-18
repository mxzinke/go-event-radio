package radio

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Radio struct {
	channelList []*Channel
}

/* Event holds the basic data structure which can be passed through the channels */
type Event struct {
	name      string
	createdAt time.Time
	payload   interface{}
}

// Creates a new Radio
// A Radio can have multiple channels in a tree structure
func NewRadio() (*Radio, error) {
	return new(Radio), nil
}

// Create a new Channel for events.
// If non is existing, it will create a new channel with parent channels.
// If one is existing, it will give you the existing channel.
func (r *Radio) Channel(path string) *Channel {
	existing := findChannelPath(path, r.channelList)
	if existing != nil {
		return existing
	}

	parentChannelPath := getParentChannelPath(path)
	var parentChannel = findChannelPath(parentChannelPath, r.channelList)
	if parentChannel == nil && parentChannelPath != "" {
		// If the parent is not existing yet
		parentChannel = r.Channel(parentChannelPath)
	}

	channels := strings.Split(path, DefaultPathSeparator)
	channel := &Channel{
		parent: parentChannel,
		name:   channels[len(channels)-1],
	}

	if parentChannel != nil {
		parentChannel.addChildren(channel)
	} else {
		r.registerMainChannel(channel)
	}

	return channel
}

// INTERNAL

func getParentChannelPath(channelPath string) string {
	splicedPath := strings.Split(channelPath, DefaultPathSeparator)
	if len(splicedPath) <= 1 {
		return ""
	}

	return strings.Join(splicedPath[:(len(splicedPath)-1)], DefaultPathSeparator)
}
func findChannelPath(path string, channelList []*Channel) *Channel {
	if len(channelList) == 0 {
		return nil
	}

	channelNames := strings.Split(path, DefaultPathSeparator)

	var foundChannel *Channel
	for _, channel := range channelList {
		if channel.name == channelNames[0] {
			if len(channelNames) == 1 {
				foundChannel = channel
				break
			}

			restOfPath := strings.Join(channelNames[1:len(channelNames)-1], DefaultPathSeparator)
			foundChannel = findChannelPath(restOfPath, channel.children)
			break
		}
	}

	return foundChannel
}

func (r *Radio) registerMainChannel(newChannel *Channel) {
	r.channelList = append(r.channelList, newChannel)
}

func (c *Channel) addChildren(newChildren *Channel) {
	if c.name == newChildren.name {
		// Internally, that the same channel can never be added as a children to itself
		log.Fatalln(fmt.Sprintf("You can't add the same channel to your own channel"))
	}

	c.children = append(c.children, newChildren)
}
