// The enum package provides a minimal version of an enum type, in a way similar
// to what is found in Java.  Each enum is a separate type and so can prevent
// type errors, ie. two enums that contain the same element will not be equal.
package util

import (
	"errors"
	"fmt"
	"reflect"
)

// https://gist.github.com/skarllot/102a5e5ea73861ff5afe
// http://stackoverflow.com/a/28393465/3728147

// Enumbase defines the operations available within an enum.
type Enumbase interface {
	// Initialize a newly created enum with the supplied values.
	// keys is a list of elements that are legal values in the enum.
	// current is the one element that exists in keys that is the enum's current
	// value.
	Init(keys []int, current int)
	// Clone creates a new enum with the same keys and current value as this.
	Clone() *EnumbaseImpl
	// Set changes the current value of the enum to a new value that is within keys.
	Set(current int) error
	// Equal tests when this enum equals the suppied other enum.  The keys and
	// the current value must match.
	Equal(other Enumbase) bool
	// Print prints to standard output the content of this enum.
	Print()
}

type EnumbaseImpl struct {
	keys       map[int]interface{}
	currentkey int
}

// this needs to be duplicated for each enum type
func NewEnumbase(keys []int, current int) *EnumbaseImpl {
	enum := new(EnumbaseImpl)
	enum.Init(keys, current)
	return enum
}

func (enum *EnumbaseImpl) Init(keys []int, current int) {
	enum.keys = make(map[int]interface{})
	for _, val := range keys {
		enum.keys[val] = nil
	}
	enum.currentkey = current
}

func (enum *EnumbaseImpl) Clone() *EnumbaseImpl {
	// http://stackoverflow.com/a/27848197/3728147
	keys := make([]int, len(enum.keys))
	i := 0
	for k := range enum.keys {
		keys[i] = k
		i++
	}
	newenum := NewEnumbase(keys, enum.currentkey)
	return newenum
}

func (enum *EnumbaseImpl) Set(current int) error {
	_, ok := enum.keys[current]
	if !ok {
		return errors.New("invalid enum value")
	} else {
		enum.currentkey = current
		return nil
	}
}

func (enum *EnumbaseImpl) Equal(other Enumbase) bool {
	otherbase := other.(*EnumbaseImpl)
	// http://stackoverflow.com/a/27160765/3728147
	typesame := reflect.TypeOf(enum) == reflect.TypeOf(otherbase)
	valuesame := enum.currentkey == otherbase.currentkey
	return typesame && valuesame
}

func (enum *EnumbaseImpl) Print() {
	fmt.Printf("%v\n", enum)
}
