package acceptor
import (
	"util"
	"fmt"
	"reflect"
	"sync"
)
type AcceptorImp struct {
	Listener    chan *util.Comm
	// Current ballot
	ballotNumer int
	// Lock for ballot (probably can get away with atomic but leave that for a later time
	ballotMutex sync.RWMutex
	pvalues     []util.Pvalue
	pvalueSet   map[string]struct{}
}

func (a *AcceptorImp) Start() {
	go func() {
		for {
			select {
			case message := <-a.Listener:
				if message.FunctionExecutor == util.ProposeFunc {
					a.Propose(message.Payload, message.ResponseChan)
				} else if message.FunctionExecutor == util.AdoptFunc {
					a.Adopt(message.Payload, message.ResponseChan)
				} else if message.FunctionExecutor == util.AcceptFunc {
					a.Accept(message.Payload, message.ResponseChan)
				} else {
					message.ResponseChan <- "Error. No such function"
				}
			}
		}
	}()
}

func NewAcceptor() *AcceptorImp {
	return &AcceptorImp{Listener: make(chan *util.Comm, 100),
		ballotNumer: 0,
		pvalues : make([]util.Pvalue, 0),
		pvalueSet: make(map[string]struct{}),
	}
}

func (a *AcceptorImp) Adopt(payload interface{}, responseChan chan interface{}) {
	adoptRequest := reflect.ValueOf(payload).Interface().(*util.Adopt)
	adoptBallot := adoptRequest.Ballot
	a.ballotMutex.Lock()
	defer a.ballotMutex.Unlock()
	if adoptBallot > a.ballotNumer {
		a.ballotNumer = adoptBallot
	}
	responseChan <- &util.AdoptResponse{a.ballotNumer, a.pvalues}
}

func (a *AcceptorImp) Accept(payload interface{}, responseChan chan interface{}) {
	acceptRequest := reflect.ValueOf(payload).Interface().(*util.Accept)
	acceptBallot := acceptRequest.Ballot
	a.ballotMutex.Lock()
	a.ballotMutex.Unlock()
	if a.ballotNumer == acceptBallot {
		proposedPValue := util.Pvalue{a.ballotNumer, acceptRequest.Slot, acceptRequest.Command}
		key := proposedPValue.Key()
		if _, ok := a.pvalueSet[key]; !ok {
			a.pvalueSet[key] = struct {}{}
			a.pvalues = append(a.pvalues, proposedPValue)
		} else {
			fmt.Printf("Error")
		}
	}
	responseChan <- &util.AcceptResponse{a.ballotNumer}
}

// This is merely used for testing the wiring; Can be ignored
func (a *AcceptorImp) Propose(payload interface{}, responseChan chan interface{}) {
	proposal := reflect.ValueOf(payload).Interface().(*util.LeaderProposal)
	fmt.Printf("A: A <- L: Ballot: %d, Slot: %d Command: %s\n", proposal.Ballot, proposal.Slot, proposal.Command)
	responseChan <- "Acceptor success"
}

func (l *AcceptorImp) Send(message *util.Comm) {
	l.Listener <- message
}



