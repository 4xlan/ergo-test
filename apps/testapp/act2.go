package testapp

import (
	"demo"
	"ergo.services/ergo/act"
	"ergo.services/ergo/gen"
	"fmt"
	"time"
)

func factory_Act2() gen.ProcessBehavior {
	return &Act2{}
}

type Act2 struct {
	act.Actor
}

// Init invoked on a start this process.
func (a *Act2) Init(args ...any) error {
	a.Log().Info("started process with name %s and args %v", a.Name(), args)
	return nil
}

//
// Methods below are optional, so you can remove those that aren't be used
//

// HandleMessage invoked if Actor received a message sent with gen.Process.Send(...).
// Non-nil value of the returning error will cause termination of this process.
// To stop this process normally, return gen.TerminateReasonNormal
// or any other for abnormal termination.
func (a *Act2) HandleMessage(from gen.PID, message any) error {
	a.Log().Info("act2: %s", a.State())
	time.Sleep(time.Second * 1)
	return nil
}

// HandleCall invoked if Actor got a synchronous request made with gen.Process.Call(...).
// Return nil as a result to handle this request asynchronously and
// to provide the result later using the gen.Process.SendResponse(...) method.
func (a *Act2) HandleCall(from gen.PID, ref gen.Ref, request any) (any, error) {
	a.Log().Info("got request from %s with reference %s", from, ref)
	switch req := request.(type) {
	case demo.Req:
		ans, err := a.Call("act3", demo.Req{
			Msg: fmt.Sprintf("%s -> act2", req.Msg),
		})
		if err != nil {
			return demo.Ans{Ok: false, Msg: req.Msg}, err
		} else {
			return ans, nil
		}
	default:
		return gen.Atom("pong"), nil
	}
}

// Terminate invoked on a termination process
func (a *Act2) Terminate(reason error) {
	a.Log().Info("terminated with reason: %s", reason)
}

// HandleMessageName invoked if split handling was enabled using SetSplitHandle(true)
// and message has been sent by name
func (a *Act2) HandleMessageName(name gen.Atom, from gen.PID, message any) error {
	return nil
}

// HandleMessageAlias invoked if split handling was enabled using SetSplitHandle(true)
// and message has been sent by alias
func (a *Act2) HandleMessageAlias(alias gen.Alias, from gen.PID, message any) error {
	return nil
}

// HandleCallName invoked if split handling was enabled using SetSplitHandle(true)
// and request was made by name
func (a *Act2) HandleCallName(name gen.Atom, from gen.PID, ref gen.Ref, request any) (any, error) {
	return gen.Atom("pong"), nil
}

// HandleCallAlias invoked if split handling was enabled using SetSplitHandle(true)
// and request was made by alias
func (a *Act2) HandleCallAlias(alias gen.Alias, from gen.PID, ref gen.Ref, request any) (any, error) {
	return gen.Atom("pong"), nil
}

// HandleLog invoked on a log message if this process was added as a logger.
// See https://docs.ergo.services/basics/logging for more information
func (a *Act2) HandleLog(message gen.MessageLog) error {
	return nil
}

// HandleEvent invoked on an event message if this process got subscribed on
// this event using gen.Process.LinkEvent or gen.Process.MonitorEvent
// See https://docs.ergo.services/basics/events for more information
func (a *Act2) HandleEvent(message gen.MessageEvent) error {
	return nil
}

// HandleInspect invoked on the request made with gen.Process.Inspect(...)
func (a *Act2) HandleInspect(from gen.PID, item ...string) map[string]string {
	a.Log().Info("got inspect request from %s", from)
	return nil
}
