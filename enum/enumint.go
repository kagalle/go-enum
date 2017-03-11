package util

// https://play.golang.org/p/4FkNSiUDMg
import (
	"errors"
	"fmt"
)

type Enumint struct {
	valuesMap map[int]interface{}
	current   int
}

func NewEnumint(valuesMap []int, current int) *Enumint {
	enumint := new(Enumint)
	enumint.Init(valuesMap, current)
	return enumint
}

// this needs to be duplicated for each enum type
func (enum *Enumint) Init(valuesMap []int, current int) {
	// enum := new(Enumbase)
	enum.valuesMap = make(map[int]interface{})
	for _, val := range valuesMap {
		enum.valuesMap[val] = nil
	}
	enum.current = current
}

func (enum *Enumint) Clone() Enum {
	newenum := new(Enumint)
	newenum.valuesMap = enum.valuesMap
	newenum.current = enum.current
	return newenum
}

func (enum *Enumint) Set(current int) error {
	if !enum.ValueExists(current) {
		return errors.New("invalid enum value")
	} else {
		enum.current = current
		return nil
	}
}

// Is the content of this Enumbase equivelant to the other Enumbase?
// Does not account for the type equality of the objects they're in.
func (enum *Enumint) Equal(other Enum) bool {
	valuesMapEqual := true
	othervaluesMap := other.getValuesMap()
	if len(enum.valuesMap) == len(othervaluesMap) {
		for index, _ := range enum.valuesMap {
			_, ok := othervaluesMap[index]
			if !ok {
				valuesMapEqual = false
				break
			}
		}
	} else {
		valuesMapEqual = false
	}
	currentEqual := enum.getCurrent() == other.getCurrent()
	return valuesMapEqual && currentEqual
}

func (enum *Enumint) getValuesMap() map[int]interface{} {
	return enum.valuesMap
}

func (enum *Enumint) getCurrent() int {
	return enum.current
}

func (enum *Enumint) ValueExists(current int) bool {
	_, ok := enum.valuesMap[current]
	return ok
}

func (enum *Enumint) Print(s string) {
	fmt.Printf("%s\t%v\n", s, enum)
}
