package radio

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

// Separator, by which the path gets spliced internally
const DefaultPathSeparator string = "."

/* A wrapper around Integer to represent Priority as an value (your value between 0 and 1000 */
type Priority uint16

// Some default priority for your usage
const (
	MIN    Priority = 0
	LOW    Priority = 100
	NORMAL Priority = 500
	HIGH   Priority = 900
	MAX    Priority = 1000
)

/* Represent a Channel where you can subscribe to. */
type Channel struct {
	parent     *Channel
	path       string
	listeners  []*eventListener
	dispatcher eventDispatcher
	children   []*Channel
}

/* Listener is function which is executed when a eventName is dispatched through a channel */
type Listener func(event Event)

/* Event holds the basic data structure which can be passed through the channels */
type Event struct {
	name      string
	createdAt time.Time
	payload   interface{}
}

var channelList []*Channel

func FindChannel(path string) *Channel {
	var foundChannel *Channel
	for _, channel := range channelList {
		if channel.path == path {
			foundChannel = channel
			break
		}
	}

	return foundChannel
}

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
		parent:     parentChannel,
		path:       path,
		dispatcher: func(event Event) {},
	}

	registerChannel(channel)

	if parentChannel != nil {
		parentChannel.addChildren(channel)
	}

	return channel, nil
}

// Returns the full path of the channel
func (c *Channel) GetPath() string {
	return c.path
}

// Gets the parent channel of the current channel
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
func getParentChannelPath(channelPath string) string {
	splicedPath := strings.Split(channelPath, DefaultPathSeparator)
	if len(splicedPath) <= 1 {
		return ""
	}

	return strings.Join(splicedPath[:(len(splicedPath)-1)], DefaultPathSeparator)
}
func findParentChannel(channelPath string) *Channel {
	parentPath := getParentChannelPath(channelPath)
	return FindChannel(parentPath)
}
func isChannelExisting(channelPath string) bool {
	return FindChannel(channelPath) != nil
}
