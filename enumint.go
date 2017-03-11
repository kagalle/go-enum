package util

// https://play.golang.org/p/4FkNSiUDMg
import (
	"errors"
	"fmt"
)

type Enumint struct {
	keys    map[int]interface{}
	current int
}

// this needs to be duplicated for each enum type
func (enum *Enumint) Init(keys []int, current int) {
	// enum := new(Enumbase)
	enum.keys = make(map[int]interface{})
	for _, val := range keys {
		enum.keys[val] = nil
	}
	enum.current = current
}

func (enum *Enumint) Clone() Enum {
	// http://stackoverflow.com/a/27848197/3728147
	// keys := make([]int, len(enum.keys))
	// for _, val := range enum.keys {
	// 	keys[val] = nil
	// }
	newenum := new(Enumint)
	newenum.keys = enum.keys
	newenum.current = enum.current
	return newenum
}

func (enum *Enumint) Set(current int) error {
	if !enum.keyExists(current) {
		return errors.New("invalid enum value")
	} else {
		enum.current = current
		return nil
	}
}

// Is the content of this Enumbase equivelant to the other Enumbase?
// Does not account for the type equality of the objects they're in.
func (enum *Enumint) Equal(other Enum) bool {
	keysEqual := true
	otherkeys := other.getKeys()
	if len(enum.keys) == len(otherkeys) {
		for index, _ := range enum.keys {
			_, ok := otherkeys[index]
			if !ok {
				keysEqual = false
				break
			}
		}
	} else {
		keysEqual = false
	}
	currentEqual := enum.getCurrent() == other.getCurrent()
	return keysEqual && currentEqual
}

func (enum *Enumint) getKeys() map[int]interface{} {
	return enum.keys
}

func (enum *Enumint) getCurrent() int {
	return enum.current
}

func (enum *Enumint) keyExists(current int) bool {
	_, ok := enum.keys[current]
	return ok
}

func (enum *Enumint) Print(s string) {
	fmt.Printf("%s\t%v\n", s, enum)
}
