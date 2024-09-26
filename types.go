package demo

import (
	"ergo.services/ergo/gen"
	"ergo.services/ergo/net/edf"
	"errors"
)

type Req struct {
	Msg string
}
type Ans struct {
	Ok  bool
	Msg string
}

type InitMsgReq struct {
	ReqLink any
}

func init() {
	types := []any{
		Req{},
		Ans{},
		InitMsgReq{},
	}

	for _, t := range types {
		err := edf.RegisterTypeOf(t)
		if err == nil || errors.Is(err, gen.ErrTaken) {
			continue
		}
		panic(err)
	}
}
