package util

const (
	Propose = "propose"
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


// Payload for replicas to propose
type ReplicaProposal struct {
	Slot int
	Command string
}

type LeaderProposal struct {
	*ReplicaProposal
	Ballot int
}