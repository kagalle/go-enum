// The enum package provides a minimal version of an enum type, in a way similar
// to what is found in Java.  Each enum is a separate type and so can prevent
// type errors, ie. two enums that contain the same element will not be equal.
package util

// https://gist.github.com/skarllot/102a5e5ea73861ff5afe
// http://stackoverflow.com/a/28393465/3728147

// Enumbase defines the operations available within an enum.
/*
type Enum interface {
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
*/
type Enum interface {
	KeyExists(current int) bool
	Equal(other *Enumbase) bool
}

type Enumbase struct {
	keys map[int]interface{}
}

// this needs to be duplicated for each enum type
func (enum *Enumbase) Init(keys []int) *Enumbase {
	// enum := new(Enumbase)
	enum.keys = make(map[int]interface{})
	for _, val := range keys {
		enum.keys[val] = nil
	}
	return enum
}

func (enum *Enumbase) KeyExists(current int) bool {
	_, ok := enum.keys[current]
	return ok
}

/*
func (enum *Enumbase) Clone() *Enumbase {
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
*/
/*
func (enum *Enumbase) Set(current int) error {
	_, ok := enum.keys[current]
	if !ok {
		return errors.New("invalid enum value")
	} else {
		enum.currentkey = current
		return nil
	}
}
*/
// Is the content of this Enumbase equivelant to the other Enumbase?
// Does not account for the type equality of the objects they're in.
func (enum *Enumbase) Equal(other *Enumbase) bool {
	isequal := true
	if len(enum.keys) == len(other.keys) {
		for _, val := range enum.keys {
			_, ok := other.keys[val.(int)]
			if !ok {
				isequal = false
				break
			}
		}
	} else {
		isequal = false
	}
	return isequal
}

/*
func (enum *Enumbase) Print() {
	fmt.Printf("%v\n", enum)
}
*/
