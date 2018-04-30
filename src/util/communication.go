package util

import "strconv"

const (
	ProposeFunc = "propose"
	AdoptFunc = "adopt"
	AcceptFunc = "accept"
)

type Comm struct {
	Payload          interface{}
	FunctionExecutor string
	ResponseChan     chan interface{}
}

type DummyPayload struct {
	Message string
}

type listener struct {
	message *Comm
}

type Sender interface {
	Send(message *Comm)
}

// This section is akin to proto definitions

// Prepare request sent from the leader also known as P1A (in paxos made complex)
type Adopt struct {
	Ballot int
}

type AdoptResponse struct {
	Ballot int
	Pvalues []Pvalue
}

type Accept struct {
	Ballot int
	Slot int
	Command string
}

type AcceptResponse struct {
	Ballot int
}

type Pvalue Accept

func (p Pvalue) Key() string {
	return strconv.Itoa(p.Ballot) + strconv.Itoa(p.Slot) + p.Command
}


// This section is used to testing the wiring can be ignored.
// Payload for replicas to propose
type ReplicaProposal struct {
	Slot int
	Command string
}

type LeaderProposal struct {
	*ReplicaProposal
	Ballot int
}


