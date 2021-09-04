package types

import "log"

type Config struct {
	Logger                                            *log.Logger
	Host, WSServerPort, WSServerPath, BroadcasterPort string
	KeyboardChan, MouseChan                           chan []byte
}
