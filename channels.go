package radio

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

// Returns the full path of the channel
func (c *Channel) GetPath() string {
	return c.path
}

// Gets the parent channel of the current channel
func (c *Channel) GetParent() *Channel {
	return c.parent
}
