package peersmsg

type MessageRaw struct {
	Data []byte
}

func (m MessageRaw) Bytes() []byte {
	return m.Data
}

func (m MessageRaw) String() string {
	return string(m.Data)
}
