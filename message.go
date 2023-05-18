package peersmsg

type Message interface {
	Bytes() []byte
	String() string
}
