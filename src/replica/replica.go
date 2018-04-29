package replica

import (
	"util"
	"fmt"
)



type ReplicaImp struct {
	slotIn int
	slotOut int
	proposals []util.ReplicaProposal
	decisions []util.ReplicaProposal
	leaders []util.Sender
}



func NewReplica(leaders []util.Sender) *ReplicaImp {
	return &ReplicaImp{
		proposals: make([]util.ReplicaProposal, 0),
		decisions: make([]util.ReplicaProposal, 0),
		leaders: leaders,
	}
}

func (r *ReplicaImp) Start() {
	// Propose a command
	proposal := &util.ReplicaProposal{1, "k:v"}
	// Response channel
	response := make(chan interface{}, 100)
	fmt.Printf("R: L <- R: %+v\n", proposal)
	r.leaders[0].Send(&util.Comm{proposal, util.Propose, response})
	fmt.Printf("\n R <- L: %+v\n", <- response)
}