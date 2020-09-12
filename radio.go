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
		return nil, errors.New(fmt.Sprintf("The channel with the name %s is already existing, try find to it instead.", path))
	}

	parentChannelPath := getParentChannelPath(path)
	var parentChannel = r.FindChannel(parentChannelPath)
	if parentChannel == nil && parentChannelPath != "" {
		// If the parent is not existing yet
		parentChannel, _ = r.NewChannel(parentChannelPath)
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

	return channel, nil
}

// FindChannel does helps you finding the right Channel by it's name
func (r *Radio) FindChannel(path string) *Channel {
	return findChannelPath(path, r.channelList)
}

// PRIVATE FUNC's

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
