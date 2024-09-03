package gate

import (
	"dgkang/demo/server/game"
	"dgkang/demo/server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
}
