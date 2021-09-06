package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type MouseAction uint8

const (
	MouseMoveAction MouseAction = iota
)

type MouseEventMessage struct {
	Action MouseAction
	DX     float64
	DY     float64
}

func NewMouseEventMessage(dx, dy float64) *MouseEventMessage {
	return &MouseEventMessage{
		Action: MouseMoveAction,
		DX:     dx,
		DY:     dy,
	}
}

func (msg *MouseEventMessage) Scale(x, y float64) *MouseEventMessage {
	msg.DX = msg.DX / x
	msg.DY = msg.DY / x

	return msg
}

func FromBytes(bz []byte, m *MouseEventMessage) error {
	bin_buf := bytes.NewBuffer(bz)
	if err := binary.Read(bin_buf, binary.BigEndian, m); err != nil {
		return fmt.Errorf("binary read: %v", err)
	}
	return nil
}

func (m *MouseEventMessage) Bytes() ([]byte, error) {
	var bin_buf bytes.Buffer
	if err := binary.Write(&bin_buf, binary.BigEndian, m); err != nil {
		return nil, fmt.Errorf("binary write: %v", err)
	}
	return bin_buf.Bytes(), nil
}
