package events

import (
	"github.com/adntgv/hidteleport/types"
	hook "github.com/robotn/gohook"
	"go.uber.org/zap"
)

type Transformer struct {
	mousePosition *types.Coordinates
	screen        *types.Screen
	logger        *zap.Logger
}

func NewTransformer(logger *zap.Logger, mouseInitialPosition *types.Coordinates, screen *types.Screen) *Transformer {
	return &Transformer{
		mousePosition: mouseInitialPosition,
		screen:        screen,
		logger:        logger.Named("transformer"),
	}
}

func (t *Transformer) Transform(ev *hook.Event) ([]byte, error, device) {
	switch ev.Kind {
	case hook.MouseUp, hook.MouseHold, hook.MouseDown, hook.MouseMove, hook.MouseDrag, hook.MouseWheel:
		bz, err := t.mouseTransform(ev)
		return bz, err, mouse
	case hook.KeyDown, hook.KeyHold, hook.KeyUp:
		bz, err := keyboardTransform(ev)
		return bz, err, keyboard
	default:
		return nil, nil, unknown
	}
}

func (t *Transformer) mouseTransform(ev *hook.Event) ([]byte, error) {
	switch ev.Kind {
	case hook.MouseMove:
		newPosition := types.Coordinates{
			X: int64(ev.X),
			Y: int64(ev.Y),
		}

		t.logger.Sugar().Debugf("new position %+v", newPosition)

		msg := types.NewMouseEventMessage(
			float64(newPosition.X-t.mousePosition.X),
			float64(newPosition.Y-t.mousePosition.Y),
		).Scale(
			float64(t.screen.Width),
			float64(t.screen.Height),
		)

		t.logger.Sugar().Debugf("msg %+v", msg)

		t.mousePosition.X = newPosition.X
		t.mousePosition.Y = newPosition.Y
		return msg.Bytes()
	}
	return nil, nil
}

func keyboardTransform(ev *hook.Event) ([]byte, error) {
	return nil, nil
}
