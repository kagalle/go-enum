package util

import "errors"

// type enummap map[string]int

type Enumbase struct {
	// instance string
	keys       map[int]interface{}
	currentkey int
}

// this needs to be duplicated for each enum type
func NewEnumbase(keys []int, current int) *Enumbase {
	enum := new(Enumbase)
	enum.Init(keys, current)
	return enum
}

func (enum *Enumbase) Init(keys []int, current int) *Enumbase {
	enum.keys = make(map[int]interface{})
	for _, val := range keys {
		enum.keys[val] = nil
	}
	enum.currentkey = current
	return enum
}

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

func (enum *Enumbase) Set(current int) (*Enumbase, error) {
	_, ok := enum.keys[current]
	if !ok {
		return nil, errors.New("invalid enum value")
	} else {
		enum.currentkey = current
		return enum, nil
	}
}
