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
