package main
import (
	"acceptor"
	"util"
	"fmt"
)

func main() {

	a1 := acceptor.NewAcceptor()
	response := make(chan interface{}, 100)
	a1.Start()
	payload := &acceptor.AcceptorPayload{1, 1, "cmd"}
	a1.Listener <- &util.Comm{payload, response}
	fmt.Printf("Response received: %s", <-response)
}
