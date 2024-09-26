package interfaces

import (
	"demo"
	"ergo.services/ergo/gen"
)

type Actor1 interface {
	ProxyCall(to any, message any) (any, error)
	Log() gen.Log
	SetSvc1(Service1)
}

type Service1 interface {
	Init(gen.Node) error
	SetMsg(Actor1)
	DoSomething(ans demo.Ans) (demo.Ans, error)
}
