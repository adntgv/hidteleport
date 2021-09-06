package emulator

import (
	"github.com/adntgv/hidteleport/types"
	"github.com/go-vgo/robotgo"
)

func GetScreenSize() *types.Screen {
	x, y := robotgo.GetScreenSize()
	return &types.Screen{
		Height: x,
		Width:  y,
	}
}
