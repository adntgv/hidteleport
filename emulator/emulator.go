package emulator

import (
	"log"
	"sync"

	"github.com/adntgv/hidteleport/types"
)

type Emulator struct {
	logger *log.Logger
	Mouse  *Mouse
}

func NewEmulator(logger *log.Logger, screen *types.Screen, mouseInChan chan []byte) *Emulator {
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
			e.logger.Println(err)
		}
	}(wg)
}
