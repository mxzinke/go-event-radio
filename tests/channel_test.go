package radio

import (
	"github.com/mxzinke/radio"
	"github.com/stretchr/testify/assert"
	"testing"
)

/* Should return new channel object when creating new channel */
func TestNewChannel(t *testing.T) {
	channelName := "test-channel.new"

	channel, err := radio.NewChannel(channelName)

	assert.NoError(t, err, "Should not return error")
	assert.IsType(t, &radio.Channel{}, channel, "Should return channel object pointer")
}

/* Should return first channel without error, but second channel should return error */
func TestNewChannelDuplicate(t *testing.T) {
	channelName := "test-channel.duplicate"

	channel1, err1 := radio.NewChannel(channelName)
	channel2, err2 := radio.NewChannel(channelName)

	/* -- first normal channel -- */
	assert.NoError(t, err1, "Should not return error")
	assert.IsType(t, &radio.Channel{}, channel1, "Should return channel object pointer")
	/* -- second channel as duplicate -- */
	assert.Error(t, err2, "Should return error")
	assert.Contains(t, err2.Error(), channelName, "Should return error containing duplicated channel name")
	assert.Nil(t, channel2, "Should not return channel object (nil-pointer)")
}

func TestGetChannelName(t *testing.T) {
	channelName := "test-channel.name"
	channel, err := radio.NewChannel(channelName)

	assert.NoError(t, err, "Should not return error")
	assert.Equal(t, channelName, channel.GetName(), "Should return channel name as a string")
}
