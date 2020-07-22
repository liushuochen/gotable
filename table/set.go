package table

import "fmt"

type Set struct {
	base []string
}

func (set *Set) Len() int {
	return len(set.base)
}

func (set *Set) Cap() int {
	return cap(set.base)
}

func (set *Set) Exit(element string) bool {
	flag := set.exit(element)
	if flag == -1 {
		return false
	}
	return true
}

func (set *Set) exit(element string) int {
	for index, data := range set.base {
		if data == element {
			return index
		}
	}

	return -1
}

func (set *Set) Clear() {
	set.base = make([]string, 0)
}

func (set *Set) Add(element string) error {
	if set.Exit(element) {
		return fmt.Errorf("value %s has exit", element)
	}
	set.base = append(set.base, element)
	return nil
}

func (set *Set) Remove(element string) error {
	position := set.exit(element)
	if position == -1 {
		return fmt.Errorf("value %s has not exit", element)
	}

	set.base = append(set.base[:position], set.base[position+1:]...)
	return nil
}