package radio_test

import (
	"github.com/mxzinke/radio"
	"github.com/stretchr/testify/assert"
	"testing"
)

/* Should return new channel object when creating new channel */
func TestNewChannel(t *testing.T) {
	testRadio, _ := radio.NewRadio()
	channelName := "test-channel.new"

	channel, err := testRadio.NewChannel(channelName)

	assert.NoError(t, err, "Should not return error")
	assert.IsType(t, &radio.Channel{}, channel, "Should return channel object pointer")
}

/* Should return first channel without error, but second channel should return error */
func TestNewChannelDuplicate(t *testing.T) {
	testRadio, _ := radio.NewRadio()
	channelPath := "test-channel.duplicate"

	channel1, err1 := testRadio.NewChannel(channelPath)
	channel2, err2 := testRadio.NewChannel(channelPath)

	/* -- first normal channel -- */
	assert.NoError(t, err1, "Should not return error")
	assert.IsType(t, &radio.Channel{}, channel1, "Should return channel object pointer")
	/* -- second channel as duplicate -- */
	assert.Error(t, err2, "Should return error")
	assert.Contains(t, err2.Error(), channelPath, "Should return error containing duplicated channel path")
	assert.Nil(t, channel2, "Should not return channel object (nil-pointer)")
}

/* Should return the correct channel name / path */
func TestGetChannelName(t *testing.T) {
	testRadio, _ := radio.NewRadio()
	channelName := "test-channel.name"

	channel, err := testRadio.NewChannel(channelName)
	actualChannelName := channel.GetPath()

	assert.NoError(t, err, "Should not return error")
	assert.Equal(t, actualChannelName, channelName, "Should return channel name as a string")
}

/* Should return the correct channel */
func TestFindChannel(t *testing.T) {
	testRadio, _ := radio.NewRadio()
	channelName := "test-channel.name"
	_, err := testRadio.NewChannel(channelName)

	actualChannel := testRadio.FindChannel(channelName)

	assert.NoError(t, err, "Should not return error")
	assert.IsType(t, &radio.Channel{}, actualChannel, "Should be type of Channel")
	assert.NotNil(t, actualChannel, "Should be type of Channel")
	assert.Equal(t, actualChannel.GetPath(), channelName, "Should return channel name as a string")
}

/* Should return no channel */
func TestFindNoChannel(t *testing.T) {
	testRadio, _ := radio.NewRadio()
	channelName := "test-channel.name"

	actualChannel := testRadio.FindChannel(channelName)

	assert.Nil(t, actualChannel, "Should be nil")
}

/* Should return correct parent channels (after creating it recursively) */
func TestGetChannelParentRecursively(t *testing.T) {
	testRadio, _ := radio.NewRadio()
	channelName := "organization.system.subsystem.parent.channel"
	parentChannelName := "organization.system.subsystem.parent"
	subsystemChannelName := "organization.system.subsystem"

	channel, err := testRadio.NewChannel(channelName)

	assert.NoError(t, err, "Should not return error")
	assert.IsType(t, &radio.Channel{}, channel, "Should return channel object pointer")
	parentChannel := channel.GetParent()
	assert.IsType(t, &radio.Channel{}, parentChannel, "Should return channel object pointer for parent")
	assert.Equal(t, parentChannelName, parentChannel.GetPath(), "Should return channel name of parent")
	subsystemChannel := parentChannel.GetParent()
	assert.IsType(t, &radio.Channel{}, subsystemChannel, "Should return channel object pointer for subsystem")
	assert.Equal(t, subsystemChannelName, subsystemChannel.GetPath(), "Should return channel name of subsystem")
}
