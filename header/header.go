package header

const (
	AlignCenter = iota
	AlignLeft
	AlignRight
)

type Header struct {
	Name			string
	defaultValue	string
	align			int
}

func CreateHeader(name string) *Header {
	h := &Header{
		Name: name,
		defaultValue: "",
		align: AlignCenter,
	}
	return h
}

func (h *Header) Default() string {
	return h.defaultValue
}

func (h *Header) SetDefault(value string) {
	h.defaultValue = value
}

func (h *Header) Align() int {
	return h.align
}

func (h *Header) AlignString() string {
	switch h.align {
	case AlignCenter:
		return "center"
	case AlignLeft:
		return "left"
	case AlignRight:
		return "right"
	default:
		return "unknown"
	}
}

func (h *Header) SetAlign(mode int) {
	switch mode {
	case AlignLeft:
		h.align = AlignLeft
	case AlignRight:
		h.align = AlignRight
	default:
		h.align = AlignCenter
	}
}
