package cell

import "unicode"

type Data struct {
	value	string
	length	int
}

func CreateData(value string) *Data {
	d := new(Data)
	d.value = value
	d.length = 0
	for _, c := range value {
		if isChinese(c) {
			d.length += 2
		} else {
			d.length += 1
		}
	}
	return d
}

func CreateEmptyData() *Data {
	return CreateData("")
}

func isChinese(c int32) bool {
	if unicode.Is(unicode.Han, c) {
		return true
	}

	for _, s := range chineseSymbol {
		if c == s {
			return true
		}
	}
	return false
}

func (d *Data) String() string {
	return d.value
}

func (d *Data) Length() int {
	return d.length
}
