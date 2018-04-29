package util

type Comm struct {
	Payload interface{}
	ResponseChan chan interface{}
}

type DummyPayload struct {
	Message string
}

