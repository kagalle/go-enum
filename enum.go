package util

import (
	"errors"
	"fmt"
	"reflect"
)

// type enummap map[string]int

type Enumbase interface {
	Init(keys []int, current int)
	Clone() *EnumbaseImpl
	Set(current int) error
	Equal(other Enumbase) bool
	Print()
}

type EnumbaseImpl struct {
	// instance string
	keys       map[int]interface{}
	currentkey int
}

// this needs to be duplicated for each enum type
func NewEnumbase(keys []int, current int) *EnumbaseImpl {
	enum := new(EnumbaseImpl)
	enum.Init(keys, current)
	fmt.Printf("NewEnumbase %v\n", enum)
	return enum
}

func (enum *EnumbaseImpl) Init(keys []int, current int) {
	enum.keys = make(map[int]interface{})
	for _, val := range keys {
		enum.keys[val] = nil
	}
	enum.currentkey = current
	fmt.Printf("Init %v\n", enum)
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
	typesame := reflect.TypeOf(enum) == reflect.TypeOf(otherbase)
	valuesame := enum.currentkey == otherbase.currentkey
	return typesame && valuesame
}

func (enum *EnumbaseImpl) Print() {
	fmt.Printf("Print %v\n%v\n", enum.keys, enum.currentkey)
}
