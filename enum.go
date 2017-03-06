package util

// type enummap map[string]int

type Enumbase struct {
	// instance string
	keys map[int]int
    currentkey int
}

type EnumOne struct {
	Enumbase
}

// this needs to be duplicated for each enum type
func NewEnum(keys []int, current int) *EnumOne {
	enum := new(EnumOne)
	enum.keys = make(map[int]int)
	for key, val := range keys {
		enum.keys[key] = val
	}
    // TODO store the key/index of the current value passed in
    enum.currentkey = // find key value in map that has a element value of 'current'
	return enum
}
