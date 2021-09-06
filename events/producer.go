package events

import (
	hook "github.com/robotn/gohook"
	"go.uber.org/zap"
)

type device uint

const (
	mouse device = iota
	keyboard
	unknown
)

type Producer struct {
	logger      *zap.Logger
	Transformer *Transformer
	OutChans    map[device]chan []byte
}

func NewProducer(transformer *Transformer, logger *zap.Logger, mouseChan, keyboardChan chan []byte) *Producer {
	return &Producer{
		logger:      logger.Named("producer"),
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
		if err != nil || dev == unknown {
			producer.logger.Sugar().Errorf("transform: %v", err)
			continue
		}

		ch, ok := producer.OutChans[dev]
		if ok {
			ch <- bz
		}
	}
}
