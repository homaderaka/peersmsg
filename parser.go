package peersmsg

import (
	"bufio"
	"io"
	"log"
)

type Parser interface {
	NextMessage(r io.Reader) (Message, error)
	SetValidator(validator func(Message) error)
	SetLogger(logger func(...interface{}))
}

type ParserRaw struct {
	separator rune
	validator func(Message) error
	logger    func(...interface{})
}

func (p *ParserRaw) NextMessage(r io.Reader) (m Message, err error) {
	reader := bufio.NewReader(r)
	line, err := reader.ReadString(byte(p.separator))
	if err != nil {
		if err != io.EOF {
			p.logger("read error:", err)
		}
		return
	}

	m = MessageRaw{Data: []byte(line[:len(line)-1])}

	// Call the validator, if one has been set.
	if p.validator != nil {
		if err = p.validator(m); err != nil {
			p.logger("validation error:", err)
			return
		}
	}

	return
}

func (p *ParserRaw) SetValidator(validator func(Message) error) {
	p.validator = validator
}

func (p *ParserRaw) SetLogger(logger func(...interface{})) {
	p.logger = logger
}

func NewParser(sep rune, opts ...ParserOption) (p Parser) {
	p = &ParserRaw{
		separator: sep,
		logger:    log.Println, // Default logger
	}

	for _, option := range opts {
		option(p)
	}

	return
}

type ParserOption func(p Parser)

func WithValidator(validator func(Message) error) ParserOption {
	return func(p Parser) {
		p.SetValidator(validator)
	}
}

func WithLogger(logger func(...interface{})) ParserOption {
	return func(p Parser) {
		p.SetLogger(logger)
	}
}
