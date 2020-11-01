package table

import "fmt"

type set struct {
	base []string
}

func (set *set) Len() int {
	return len(set.base)
}

func (set *set) Cap() int {
	return cap(set.base)
}

func (set *set) Exist(element string) bool {
	flag := set.exist(element)
	if flag == -1 {
		return false
	}
	return true
}

func (set *set) exist(element string) int {
	for index, data := range set.base {
		if data == element {
			return index
		}
	}

	return -1
}

func (set *set) Clear() {
	set.base = make([]string, 0)
}

func (set *set) Add(element string) error {
	if set.Exist(element) {
		return fmt.Errorf("value %s has exit", element)
	}
	set.base = append(set.base, element)
	return nil
}

func (set *set) Remove(element string) error {
	position := set.exist(element)
	if position == -1 {
		return fmt.Errorf("value %s has not exit", element)
	}

	set.base = append(set.base[:position], set.base[position+1:]...)
	return nil
}
