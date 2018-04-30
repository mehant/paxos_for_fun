package main
import (
	"acceptor"
	"leader"
	"replica"
	"util"
)

func main() {

	a1 := acceptor.NewAcceptor()
	a1.Start()
	l1 := leader.NewLeader([]util.Sender{a1})
	l1.Start()
	r1 := replica.NewReplica([]util.Sender{l1})
	r1.Propose("K: v")
}
