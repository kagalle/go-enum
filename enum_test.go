package util

// https://play.golang.org/p/4FkNSiUDMg
import (
	"fmt"
	"testing"
)

type EnumOne struct {
	Enumbase
}

// this needs to be duplicated for each enum type
func NewEnumOne(keys []int, current int) *EnumOne {
	enum := new(EnumOne)
	enum.Init(keys, current)
	return enum
}

func TestEnum(t *testing.T) {
	e := NewEnumOne([]int{1, 2, 3}, 2)
	fmt.Printf("%v\n", e)
	f := e.Clone()
	f.Set(1)
	fmt.Printf("%v\n", f)
}
