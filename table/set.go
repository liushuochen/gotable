package table

import (
	"fmt"
	"github.com/liushuochen/gotable/header"
)

type Set struct {
	base []*header.Column
}

func (set *Set) Len() int {
	return len(set.base)
}

func (set *Set) Cap() int {
	return cap(set.base)
}

func (set *Set) Exist(element string) bool {
	return set.exist(element) != -1
}

func (set *Set) exist(element string) int {
	for index, data := range set.base {
		if data.String() == element {
			return index
		}
	}

	return -1
}

func (set *Set) Clear() {
	set.base = make([]*header.Column, 0)
}

func (set *Set) Add(element string) error {
	if set.Exist(element) {
		return fmt.Errorf("value %s has exit", element)
	}

	newHeader := header.CreateColumn(element)
	set.base = append(set.base, newHeader)
	return nil
}

func (set *Set) Remove(element string) error {
	position := set.exist(element)
	if position == -1 {
		return fmt.Errorf("value %s has not exit", element)
	}

	set.base = append(set.base[:position], set.base[position+1:]...)
	return nil
}

func (set *Set) Get(name string) *header.Column {
	for _, h := range set.base {
		if h.String() == name {
			return h
		}
	}
	return nil
}

func (set *Set) Equal(other *Set) bool {
	if set.Len() != other.Len() {
		return false
	}

	// TODO (gotable 4) improvement: Use goroutine to speed up validation
	for index := range set.base {
		if !set.base[index].Equal(other.base[index]) {
			return false
		}
	}
	return true
}
