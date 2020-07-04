package radio

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

/* Represent a Channel where you can subscribe to. */
type Channel struct {
	parent    *Channel
	path      string
	listeners []*Listener
	children  []*Channel
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

// Create a new Channel to work with it
func NewChannel(path string) (*Channel, error) {
	if isChannelExisting(path) {
		return nil, errors.New(fmt.Sprintf("The channel with the path %s is already existing, try listen to it instead.", path))
	}

	parentChannelPath := getParentChannelPath(path)
	var parentChannel = findParentChannel(parentChannelPath)
	if parentChannel == nil && parentChannelPath != "" {
		parentChannel, _ = NewChannel(parentChannelPath)
	}

	channel := &Channel{
		parent: parentChannel,
		path:   path,
	}

	registerChannel(channel)

	if parentChannel != nil {
		parentChannel.addChildren(channel)
	}

	return channel, nil
}

// Returns the full path of the channel
func (c *Channel) GetPath() string {
	if c == nil {
		return ""
	}

	return c.path
}

func (c *Channel) GetParent() *Channel {
	return c.parent
}

func (c *Channel) addChildren(newChildren *Channel) {
	if c.path == newChildren.path {
		// Internally, that the same channel can never be added as a children to itself
		log.Fatalln(fmt.Sprintf("You can't add the same channel to your own channel"))
	}

	c.children = append(c.children, newChildren)
}

func registerChannel(newChannel *Channel) {
	channelList = append(channelList, newChannel)
}

func findChannel(channelPath string) *Channel {
	var foundChannel *Channel
	for _, channel := range channelList {
		if channel.path == channelPath {
			foundChannel = channel
			break
		}
	}

	return foundChannel
}

func getParentChannelPath(channelPath string) string {
	splicedPath := strings.Split(channelPath, ".")
	if len(splicedPath) <= 1 {
		return ""
	}

	return strings.Join(splicedPath[:(len(splicedPath)-1)], ".")
}
func findParentChannel(channelPath string) *Channel {
	parentPath := getParentChannelPath(channelPath)
	return findChannel(parentPath)
}

func isChannelExisting(channelPath string) bool {
	return findChannel(channelPath) != nil
}
