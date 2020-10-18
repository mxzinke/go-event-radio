package radio_test

import (
	"github.com/mxzinke/radio"
	"github.com/stretchr/testify/assert"
	"testing"
)

/* Should return new channel object when creating new channel */
func TestNewChannel(t *testing.T) {
	testRadio, _ := radio.NewRadio()
	channelName := "test-channel:new"

	channel := testRadio.Channel(channelName)

	assert.IsType(t, &radio.Channel{}, channel,
		"Should return channel object pointer")
}

/* Should return first channel without error, but second channel should return error */
func TestNewChannelDuplicate(t *testing.T) {
	testRadio, _ := radio.NewRadio()
	channelPath := "test-channel:duplicate"

	channel1 := testRadio.Channel(channelPath)
	channel2 := testRadio.Channel(channelPath)

	/* -- first normal channel -- */
	assert.IsType(t, &radio.Channel{}, channel1,
		"Should return channel object pointer")
	/* -- second channel as duplicate -- */
	assert.IsType(t, &radio.Channel{}, channel2,
		"Should return channel object pointer")
}

/* Should return the correct channel name / path */
func TestGetChannelName(t *testing.T) {
	testRadio, _ := radio.NewRadio()
	channelName := "test-channel:name"

	channel := testRadio.Channel(channelName)
	actualChannelName := channel.GetPath()

	assert.Equal(t, actualChannelName, channelName,
		"Should return channel name as a string")
}

/* Should return correct parent channels (after creating it recursively) */
func TestGetChannelParentRecursively(t *testing.T) {
	testRadio, _ := radio.NewRadio()
	channelName := "organization:system:subsystem:parent:channel"
	parentChannelName := "organization:system:subsystem:parent"
	subsystemChannelName := "organization:system:subsystem"

	channel := testRadio.Channel(channelName)

	assert.IsType(t, &radio.Channel{}, channel,
		"Should return channel object pointer")

	parentChannel := channel.GetParent()
	assert.IsType(t, &radio.Channel{}, parentChannel,
		"Should return channel object pointer for parent")
	assert.Equal(t, parentChannelName, parentChannel.GetPath(),
		"Should return channel name of parent")

	subsystemChannel := parentChannel.GetParent()
	assert.IsType(t, &radio.Channel{}, subsystemChannel,
		"Should return channel object pointer for subsystem")
	assert.Equal(t, subsystemChannelName, subsystemChannel.GetPath(),
		"Should return channel name of subsystem")
}
