package emulator

import (
	"sync"

	"github.com/adntgv/hidteleport/types"
	"go.uber.org/zap"
)

type Emulator struct {
	logger *zap.Logger
	Mouse  *Mouse
}

func NewEmulator(logger *zap.Logger, screen *types.Screen, mouseInChan chan []byte) *Emulator {
	return &Emulator{
		logger: logger,
		Mouse:  NewMouse(logger, screen, mouseInChan),
	}
}

func (e *Emulator) Run() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		if err := e.Mouse.Run(); err != nil {
			e.logger.Sugar().Error(err)
		}
	}(wg)
}
