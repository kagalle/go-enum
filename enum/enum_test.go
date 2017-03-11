package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type EnumTest struct {
	*Enumint
}

// this needs to be duplicated for each enum type
func NewEnumTest(valuesMap []int, current int) *EnumTest {
	enum := new(EnumTest)
	enum.Enumint = NewEnumint(valuesMap, current)
	return enum
}

type EnumOther struct {
	*Enumint
}

// this needs to be duplicated for each enum type
func NewEnumOther(valuesMap []int, current int) *EnumOther {
	enum := new(EnumOther)
	enum.Enumint = NewEnumint(valuesMap, current)
	return enum
}

func TestEnum(t *testing.T) {
	e := NewEnumTest([]int{1, 2, 3}, 2)
	e.Print("e2")
	assert.False(t, e.ValueExists(0))
	assert.Equal(t, e.getValuesMap()[1], nil)
	assert.Equal(t, e.getCurrent(), 2)
	f := e.Clone()
	f.Set(2)
	f.Print("f2")
	assert.True(t, e.Equal(f))
	f.Set(1)
	f.Print("f1")
	assert.False(t, e.Equal(f))

	g := NewEnumOther([]int{1, 2, 4}, 2)
	g.Print("g2")
	AllowTestOnly("e - AllowTestOnly", e)

	// ./enum_test.go:58: cannot use g (type *EnumOther) as type *EnumTest in argument to AllowTestOnly
	// AllowTestOnly("g - AllowTestOnly", g)

	assert.False(t, e.Equal(g))
	assert.False(t, g.Equal(e))
}

func AllowTestOnly(s string, e *EnumTest) {
	e.Print(s)
}
