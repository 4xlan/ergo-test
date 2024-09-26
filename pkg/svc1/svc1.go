package svc1

import (
	"demo"
	"demo/pkg/constants"
	"demo/pkg/interfaces"
	"ergo.services/ergo/gen"
	"fmt"
	"log"
	"math/rand"
	"time"
)

var _ interfaces.Service1 = &Svc1{}

type Svc1 struct {
	msg  interfaces.Actor1
	wait chan struct{}
}

func (s *Svc1) Init(msgLink gen.Node) error {
	s.wait = make(chan struct{})
	err := msgLink.Send(constants.Act1, demo.InitMsgReq{ReqLink: s})
	if err != nil {
		return err
	}
	return nil
}

func (s *Svc1) SetMsg(msg interfaces.Actor1) {
	if s.msg == nil {
		s.msg = msg
		s.msg.Log().Info("svc1 msg set, closing the channel")
		close(s.wait)
	}
}

func (s *Svc1) Start() {
	log.Println("svc1 started in sleeping state")
	<-s.wait
	for {
		time.Sleep(1 * time.Second)

		// Send call to act2
		ans, err := s.msg.ProxyCall(constants.Act2, demo.Req{Msg: "act1"})
		if err != nil {
			s.msg.Log().Error("call act2 failed: %s", err)
		} else {
			switch answer := ans.(type) {
			case demo.Ans:
				s.msg.Log().Info("chain call result: %s", answer.Msg)
			default:
				s.msg.Log().Error("unknown answer type: %T", answer)
			}
		}
	}
}

func (s *Svc1) DoSomething(ans demo.Ans) (demo.Ans, error) {
	s.msg.Log().Info("doing something")
	tmp := rand.Intn(10)
	ans.Msg = fmt.Sprintf("%d | %s", tmp, ans.Msg)
	return ans, nil
}
