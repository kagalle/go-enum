package util

// type enummap map[string]int

type Enumbase struct {
	// instance string
	keys       map[int]interface{}
	currentkey int
}

// this needs to be duplicated for each enum type
func NewEnumbase(keys []int, current int) *Enumbase {
	enum := new(EnumOne)
	enum.keys = make(map[int]interface{})
	for _, val := range keys {
		enum.keys[val] = nil
	}
	// TODO store the key/index of the current value passed in
	enum.currentkey = current
	return enum
}

type EnumOne struct {
	Enumbase
}

// this needs to be duplicated for each enum type
func NewEnumOne(keys []int, current int) *EnumOne {
	enum := NewEnumbase(keys, current)
	return enum
}
