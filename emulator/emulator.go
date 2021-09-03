package emulator

import (
	"log"
	"sync"
)

type Emulator struct {
	logger log.Logger
	Mouse  *Mouse
}

func (e *Emulator) Run() error {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		if err := e.Mouse.Run(); err != nil {
			e.logger.Println(err)
		}
	}(wg)

	return nil
}
