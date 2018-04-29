package leader

import (
	"util"
	"fmt"
	"reflect"
)


type LeaderImp struct {
	listener chan *util.Comm
	acceptors []util.Sender
}

func NewLeader(acceptors []util.Sender) *LeaderImp {
	return &LeaderImp{make(chan *util.Comm, 100), acceptors}
}

func (l *LeaderImp) Start() {
	go func() {
		for {
			select {
			case message := <- l.listener:
				functionExecutor := message.FunctionExecutor
				// Dispatch the appropriate function
				if functionExecutor == util.Propose {
					l.Propose(message.Payload, message.ResponseChan)
				} else {
					message.ResponseChan <- "ERROR: No such function"
				}
			}
		}
	}()
}

func (l *LeaderImp) Propose(payload interface{}, response chan interface{}) {
	proposal := reflect.ValueOf(payload).Interface().(*util.ReplicaProposal)
	fmt.Printf("L: L <- R. Slot: %d command: %s\n", proposal.Slot, proposal.Command)
	leaderResponse := make(chan interface{})
	leaderProposal := &util.LeaderProposal{proposal, 5}
	fmt.Printf("L: A <- L :%+v\n", leaderProposal)
	l.acceptors[0].Send(&util.Comm{leaderProposal, util.Propose, leaderResponse})
	acceptorResponse := <-leaderResponse
	fmt.Printf("L: L <- A: %s\n", acceptorResponse)
	response <- acceptorResponse
}

func (l *LeaderImp) Send(message *util.Comm) {
	l.listener <- message
}
