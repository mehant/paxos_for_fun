package acceptor
import (
	"util"
	"fmt"
	"reflect"
)
type AcceptorImp struct {
	Listener chan *util.Comm
}

type AcceptorPayload struct {
	Ballot  int
	Slot    int
	Command string
}


func (a *AcceptorImp) Start() {
	go func() {
		for {
			select {
			case message := <-a.Listener:
				if message.FunctionExecutor == util.Propose {
					a.Propose(message.Payload, message.ResponseChan)
				} else {
					message.ResponseChan <- "Error. No such function"
				}
			}
		}
	}()
}

func NewAcceptor() *AcceptorImp {
	return &AcceptorImp{make(chan *util.Comm, 100)}
}

func (a *AcceptorImp) Propose(payload interface{}, responseChan chan interface{}) {
	proposal := reflect.ValueOf(payload).Interface().(*util.LeaderProposal)
	fmt.Printf("A: A <- L: Ballot: %d, Slot: %d Command: %s\n", proposal.Ballot, proposal.Slot, proposal.Command)
	responseChan <- "Acceptor success"
}

func (l *AcceptorImp) Send(message *util.Comm) {
	l.Listener <- message
}



