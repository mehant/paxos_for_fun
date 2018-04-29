package acceptor
import (
	"util"
	"fmt"
	"reflect"
)
type acceptorImp struct {
	Listener chan *util.Comm
}

type AcceptorPayload struct {
	Ballot  int
	Slot    int
	Command string
}


func (a *acceptorImp) Start() {
	go func() {
		for {
			select {
			case message := <-a.Listener:
				a.Receive(message.Payload, message.ResponseChan)
			}
		}
	}()
}

func NewAcceptor() *acceptorImp {
	return &acceptorImp{make(chan *util.Comm, 100), make(chan struct{})}
}

func (a *acceptorImp) Receive(payload interface{}, responseChan chan interface{}) {
	v := reflect.ValueOf(payload)
	request := v.Interface().(*AcceptorPayload)
	fmt.Printf("Acceptor request: Ballot: %d, Slot: %d Command: %s\n", request.Ballot, request.Slot, request.Command)
	responseStr := fmt.Sprintf("Ballot: %d, Slot: %d Command: %s", request.Ballot, request.Slot, request.Command)
	responseChan <- responseStr
}

