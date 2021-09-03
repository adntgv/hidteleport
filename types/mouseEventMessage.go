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
	DX     int64
	DY     int64
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
