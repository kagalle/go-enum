package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type EnumTest struct {
	*EnumIntImpl
}

// this needs to be duplicated for each enum type
func NewEnumTest(keys []int, current int) *EnumTest {
	enum := new(EnumTest)
	enum.EnumIntImpl = new(EnumIntImpl)
	enum.EnumIntImpl.Init(keys, current)
	return enum
}

func (enum *EnumTest) Clone() *EnumTest {
	newenum := new(EnumTest)
	newenum.EnumIntImpl = enum.EnumIntImpl.Clone()
	return newenum
}

func (enum *EnumTest) Equal(other *EnumTest) bool {
	return enum.EnumIntImpl.Equal(other.EnumIntImpl)
}

type EnumOther struct {
	*EnumIntImpl
}

// this needs to be duplicated for each enum type
func NewEnumOther(keys []int, current int) *EnumOther {
	enum := new(EnumOther)
	enum.EnumIntImpl = new(EnumIntImpl)
	enum.EnumIntImpl.Init(keys, current)
	return enum
}
func (enum *EnumOther) Clone() *EnumOther {
	newenum := new(EnumOther)
	newenum.EnumIntImpl = enum.EnumIntImpl.Clone()
	return newenum
}

func TestEnum(t *testing.T) {
	e := new(EnumTest)
	e.Init([]int{1, 2, 3}, 2)
	e.Print()
	f := e.Clone()
	f.Set(2)
	f.Print()
	assert.True(t, e.Equal(f))
	f.Set(1)
	f.Print()
	assert.False(t, e.Equal(f))

	// g := NewEnumOther([]int{1, 2, 3}, 2)
	AllowTestOnly(e)
	//	AllowTestOnly(g)

}

func AllowTestOnly(e *EnumTest) {
	e.Print()
}
