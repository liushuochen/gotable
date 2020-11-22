package header

type Header struct {
	Name			string
	defaultValue	string
}

func (h *Header) Default() string {
	return h.defaultValue
}

func (h *Header) SetDefault(value string) {
	h.defaultValue = value
}