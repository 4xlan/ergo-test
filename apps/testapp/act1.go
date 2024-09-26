package testapp

import (
	"demo"
	"demo/pkg/interfaces"
	"demo/pkg/svc1"
	"ergo.services/ergo/act"
	"ergo.services/ergo/gen"
	"reflect"
	"sync"
)

func factory_Act1() gen.ProcessBehavior {
	return &Act1{}
}

type Act1 struct {
	act.Actor
	svc    *svc1.Svc1
	mtx    sync.Mutex
	svcSet bool
}

// Init invoked on a start this process.
func (a *Act1) Init(args ...any) error {
	a.Log().Info("started process with name %s and args %v", a.Name(), args)
	a.svcSet = false
	return nil
}

//
// Methods below are optional, so you can remove those that aren't be used
//

// HandleMessage invoked if Actor received a message sent with gen.Process.Send(...).
// Non-nil value of the returning error will cause termination of this process.
// To stop this process normally, return gen.TerminateReasonNormal
// or any other for abnormal termination.
func (a *Act1) HandleMessage(from gen.PID, message any) error {
	a.Log().Info("got message from %s", from)
	switch msg := message.(type) {
	case demo.InitMsgReq:
		a.Log().Info("got init message")
		switch link := msg.ReqLink.(type) {
		case *svc1.Svc1:
			a.Log().Info("got svc1 link")
			a.SetSvc1(link)
			a.svc.SetMsg(a)
		default:
			a.Log().Error("got unknown link %v", reflect.TypeOf(link))
		}
	default:
		a.Log().Error("got unknown message %v", message)
	}
	return nil
}

// HandleCall invoked if Actor got a synchronous request made with gen.Process.Call(...).
// Return nil as a result to handle this request asynchronously and
// to provide the result later using the gen.Process.SendResponse(...) method.
func (a *Act1) HandleCall(from gen.PID, ref gen.Ref, request any) (any, error) {
	a.Log().Info("got request from %s with reference %s", from, ref)
	switch req := request.(type) {
	case demo.Ans:
		answer, err := a.svc.DoSomething(req)
		if err != nil {
			return nil, err
		} else {
			return answer, nil
		}
	default:
		a.Log().Error("got unknown request %v", request)
	}
	return gen.Atom("pong"), nil
}

// Terminate invoked on a termination process
func (a *Act1) Terminate(reason error) {
	a.Log().Info("terminated with reason: %s", reason)
}

func (a *Act1) SetSvc1(svc interfaces.Service1) {
	a.mtx.Lock()
	defer a.mtx.Unlock()

	if !a.svcSet {
		a.svc = svc.(*svc1.Svc1)
		a.Log().Info("svc1 link set")
		a.svcSet = true
	}
}

func (a *Act1) ProxyCall(to any, message any) (any, error) {
	return a.Call(to, message)
}
