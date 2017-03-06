package util

// https://play.golang.org/p/4FkNSiUDMg
import (
	"fmt"
	"testing"
)

func TestEnum(t *testing.T) {
	e := NewEnum([]int{1, 2, 3})
	fmt.Printf("%v\n", e)
}

func DoSomething(e EnumOne) {
	fmt.Printf("%v\n", e)
}
