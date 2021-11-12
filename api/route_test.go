package api

import (
	"testing"

	"net/http"
	"net/http/httptest"

	. "gopkg.in/check.v1"
)

type MySuite struct{}

var _ = Suite(&MySuite{})

func Test(t *testing.T) { TestingT(t) }

func (m *MySuite) SetUpSuite(c *C) {
	Router()
}
func (m *MySuite) TestHandler(c *C) {
	//
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/db/mysql", nil)
	c.Assert(err, IsNil)
	c.Assert(rr.Code, Equals, http.StatusOK)
	c.Log("xxx")
	r.ServeHTTP(rr, req)
}
