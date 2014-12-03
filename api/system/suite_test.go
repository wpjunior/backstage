package system

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/albertoleal/backstage/db"
	"github.com/tsuru/config"
	"github.com/zenazn/goji/web"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type S struct {
	router   *web.Mux
	handler  http.HandlerFunc
	recorder *httptest.ResponseRecorder
}

func (s *S) SetUpSuite(c *C) {
	config.Set("database:url", "127.0.0.1:27017")
	config.Set("database:name", "backstage_api_system_test")
}

func (s *S) TearDownSuite(c *C) {
	storage, err := db.Conn()
	c.Assert(err, IsNil)
	defer storage.Close()
	config.Unset("database:url")
	config.Unset("database:name")
}

var _ = Suite(&S{})
