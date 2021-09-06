package emulator

import (
	"github.com/adntgv/hidteleport/types"
	"github.com/go-vgo/robotgo"
	"go.uber.org/zap"
)

type Mouse struct {
	InChan chan []byte
	Screen *types.Screen
	logger *zap.Logger
}

func NewMouse(logger *zap.Logger, screen *types.Screen, inChan chan []byte) *Mouse {
	return &Mouse{
		logger: logger,
		Screen: screen,
		InChan: inChan,
	}
}

func (m *Mouse) Run() error {
	for bz := range m.InChan {
		msg := new(types.MouseEventMessage)
		err := types.FromBytes(bz, msg)
		if err != nil {
			m.logger.Sugar().Error(err)
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
