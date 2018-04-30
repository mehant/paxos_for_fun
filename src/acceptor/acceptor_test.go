package acceptor

import (
	. "gopkg.in/check.v1"
	"testing"
	"util"
	"reflect"
)
func Test(t *testing.T) { TestingT(t) }

type AcceptorTestSuite struct {}

var _ = Suite(&AcceptorTestSuite{})



func (a *AcceptorTestSuite) TestAdopt(c *C) {
	// Initialize acceptor
	acceptor := NewAcceptor()
	acceptor.Start()
	adopt := &util.Adopt{1}
	response := make(chan interface{}, 100)
	acceptor.Send(&util.Comm{adopt, util.AdoptFunc, response})
	adoptResponse := reflect.ValueOf(<-response).Interface().(*util.AdoptResponse)
	// Verify that the ballot is increased
	c.Assert(adoptResponse.Ballot, Equals, 1)
}
