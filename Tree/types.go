package Tree

type Ordered interface {
	// Greater returns true if the receiver is greater than the other element.
	Greater(other Ordered) bool
	// Less returns true if the receiver is less than the other element.
	Less(other Ordered) bool
	// Equal returns true if the receiver is equal to the other element.
	Equal(other Ordered) bool
}
