package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type EnumTest struct {
	*EnumInt
}

// this needs to be duplicated for each enum type
func NewEnumTest(keys []int, current int) *EnumTest {
	enum := new(EnumTest)
	enum.EnumInt = NewEnumInt(keys, current)
	return enum
}

func (enum *EnumTest) Set(current int) error {
	return enum.EnumInt.Set(current)
}

func (enum *EnumTest) Clone() *EnumTest {
	newenum := new(EnumTest)
	newenum.EnumInt = enum.EnumInt.Clone()
	return newenum
}

func (enum *EnumTest) Equal(other *EnumTest) bool {
	return enum.EnumInt.Equal(other.EnumInt)
}

func TestEnum(t *testing.T) {
	e := NewEnumInt([]int{1, 2, 3}, 2)
	e.Print()
	f := e.Clone()
	f.Set(2)
	f.Print()
	assert.True(t, e.Equal(f))
	f.Set(1)
	f.Print()
	assert.False(t, e.Equal(f))
}
