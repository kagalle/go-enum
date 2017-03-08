package util

// https://play.golang.org/p/4FkNSiUDMg
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// type EnumInt interface {
// 	Enumbase
// }

type EnumIntImpl struct {
	EnumbaseImpl
}

// this needs to be duplicated for each enum type
func NewEnumInt(keys []int, current int) *EnumIntImpl {
	enum := new(EnumIntImpl)
	enum.Init(keys, current)
	return enum
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
