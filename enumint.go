package util

// https://play.golang.org/p/4FkNSiUDMg
import (
	"errors"
	"fmt"
)

type EnumInt interface {
	KeyExists(current int) bool
	Set(current int) error
}

type EnumIntImpl struct {
	keys    *Enumbase
	current int
}

// this needs to be duplicated for each enum type
func (enum *EnumIntImpl) Init(keys []int, current int) *EnumIntImpl {
	// enum := new(EnumIntImpl)
	enum.keys = new(Enumbase)
	enum.keys.Init(keys)
	enum.current = current
	return enum
}

func (enum *EnumIntImpl) KeyExists(current int) bool {
	return enum.keys.KeyExists(current)
}

func (enum *EnumIntImpl) Set(current int) error {
	if !enum.KeyExists(current) {
		return errors.New("invalid enum value")
	} else {
		enum.current = current
		return nil
	}
}

func (enum *EnumIntImpl) Clone() *EnumIntImpl {
	// http://stackoverflow.com/a/27848197/3728147
	newenum := new(EnumIntImpl)
	newenum.keys = enum.keys
	newenum.current = enum.current
	return newenum
}

func (enum *EnumIntImpl) Print() {
	fmt.Printf("%v\t%v\n", enum.keys, enum.current)
}

func (enum *EnumIntImpl) Equal(other *EnumIntImpl) bool {
	return (enum.keys.Equal(other.keys) && (enum.current == other.current))
}
