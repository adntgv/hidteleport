package events

import (
	"github.com/adntgv/hidteleport/types"
	hook "github.com/robotn/gohook"
)

type Transformer struct {
	mousePosition *types.Coordinates
	screen        *types.Screen
}

func NewTransformer(mouseInitialPosition *types.Coordinates, screen *types.Screen) *Transformer {
	return &Transformer{
		mousePosition: mouseInitialPosition,
		screen:        screen,
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
			X: uint64(ev.X),
			Y: uint64(ev.Y),
		}

		msg := types.NewMouseEventMessage(
			float64(newPosition.X-t.mousePosition.X),
			float64(newPosition.Y-t.mousePosition.Y),
		)

		t.mousePosition.X = newPosition.X
		t.mousePosition.Y = newPosition.Y
		return msg.Bytes()
	}
	return nil, nil
}

func keyboardTransform(ev *hook.Event) ([]byte, error) {
	return nil, nil
}
