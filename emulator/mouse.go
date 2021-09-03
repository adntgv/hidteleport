package emulator

import (
	"log"

	"github.com/adntgv/hidteleport/types"
	"github.com/go-vgo/robotgo"
)

type Screen struct {
	Width  int
	Height int
}

func (s *Screen) ComputePositionAt(x, y float64) (int, int) {
	width := float64(s.Width) * x
	height := float64(s.Height) * y
	return int(width), int(height)
}

type Mouse struct {
	InChan chan []byte
	Screen Screen
	logger log.Logger
}

func (m *Mouse) Run() error {
	for bz := range m.InChan {
		msg := new(types.MouseEventMessage)
		err := types.FromBytes(bz, msg)
		if err != nil {
			m.logger.Println(err)
			continue
		}
		m.Handle(msg)
	}

	return nil
}

func (m *Mouse) Handle(msg *types.MouseEventMessage) {
	switch msg.Action {
	case types.MouseMoveAction:
		m.MoveRelative(int(msg.DX), int(msg.DY))
	}
}

func (m *Mouse) MoveAbsolute(x, y float64) {
	robotgo.MoveMouse(m.Screen.ComputePositionAt(x, y))
}

func (m *Mouse) MoveRelative(dx, dy int) {
	robotgo.MoveRelative(dx, dy)
}
