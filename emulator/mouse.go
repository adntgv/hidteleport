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
		logger: logger.Named("mouse"),
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
		m.logger.Sugar().Debugf("received message %+v", msg)
		m.Handle(msg)
	}

	return nil
}

func (m *Mouse) Handle(msg *types.MouseEventMessage) {
	switch msg.Action {
	case types.MouseMoveAction:
		msg.Unscale(float64(m.Screen.Width), float64(m.Screen.Height))
		m.MoveRelative(int(msg.DX), int(msg.DY))
	case types.MouseClickAction:
		m.Click(msg.Button)
	}
}

func (m *Mouse) MoveAbsolute(x, y float64) {
	robotgo.MoveMouse(int(x), int(y))
}

func (m *Mouse) MoveRelative(dx, dy int) {
	robotgo.MoveRelative(dx, dy)
}

func GetMousePosition() *types.Coordinates {
	x, y := robotgo.GetMousePos()
	return &types.Coordinates{
		X: int64(x),
		Y: int64(y),
	}
}

func (m *Mouse) Click(button types.MouseButton) {
	robotgo.MouseClick(button.String())
}
