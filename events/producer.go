package events

import (
	"log"

	hook "github.com/robotn/gohook"
)

type device uint

const (
	mouse device = iota
	keyboard
	unknown
)

type Producer struct {
	logger      *log.Logger
	Transformer *Transformer
	OutChans    map[device]chan []byte
}

func NewProducer(transformer *Transformer, logger *log.Logger, mouseChan, keyboardChan chan []byte) *Producer {
	return &Producer{
		logger:      logger,
		Transformer: transformer,
		OutChans: map[device]chan []byte{
			mouse:    mouseChan,
			keyboard: keyboardChan,
		},
	}
}

func (producer *Producer) Run() {
	EvChan := hook.Start()
	defer hook.End()

	for ev := range EvChan {
		bz, err, dev := producer.Transformer.Transform(&ev)
		if err != nil {
			producer.logger.Printf("transform: %v", err)
			continue
		}
		producer.OutChans[dev] <- bz
	}
}
