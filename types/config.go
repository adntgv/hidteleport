package types

import "go.uber.org/zap"

type Config struct {
	Logger                                            *zap.Logger
	Host, WSServerPort, WSServerPath, BroadcasterPort string
	KeyboardChan, MouseChan                           chan []byte
}
