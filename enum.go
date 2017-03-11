// The enum package provides a minimal version of an enum type, in a way similar
// to what is found in Java.  Each enum is a separate type and so can prevent
// type errors, ie. two enums that contain the same element will not be equal.
package util

// https://gist.github.com/skarllot/102a5e5ea73861ff5afe
// http://stackoverflow.com/a/28393465/3728147

// Enumbase defines the operations available within an enum.

type Enum interface {
	// Initialize a newly created enum with the supplied values.
	// keys is a list of elements that are legal values in the enum.
	// current is the one element that exists in keys that is the enum's current
	// value.
	Init(keys []int, current int)
	// Clone creates a new enum with the same keys and current value as this.
	Clone() Enum
	// Set changes the current value of the enum to a new value that is within keys.
	Set(current int) error
	// Equal tests when this enum equals the suppied other enum.  The keys and
	// the current value must match.
	Equal(other Enum) bool
	// Print prints to standard output the content of this enum.
	Print(s string)
	// support for Equal()
	getKeys() map[int]interface{}
	getCurrent() int
	keyExists(current int) bool
}
