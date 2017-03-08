package util

// https://play.golang.org/p/4FkNSiUDMg
import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type EnumInt interface {
	Enumbase
}

type EnumIntImpl struct {
	EnumbaseImpl
}

// this needs to be duplicated for each enum type
func NewEnumOne(keys []int, current int) *EnumIntImpl {
	enum := new(EnumIntImpl)
	enum.Init(keys, current)
	return enum
}

func TestEnum(t *testing.T) {
	e := NewEnumOne([]int{1, 2, 3}, 2)
	fmt.Printf("TestEnum %v\n", e.currentkey)
	e.Print()
	f := e.Clone()
	f.Set(2)
	f.Print()
	assert.Equal(t, e, f)
	f.Set(1)
	assert.NotEqual(t, e, f)
}
