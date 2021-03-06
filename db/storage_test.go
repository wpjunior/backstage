package db

import (
	"testing"

	"github.com/tsuru/config"
	. "gopkg.in/check.v1"
)

//Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func (s *S) SetUpSuite(c *C) {
	config.Set("database:url", "127.0.0.1:27017")
	config.Set("database:name", "backstage_db_test")
}

func (s *S) TearDownSuite(c *C) {
	storage, err := Conn()
	c.Assert(err, IsNil)
	defer storage.Close()
	config.Unset("database:url")
	config.Unset("database:name")
	// storage.Collection("services").Database.DropDatabase()
}

func (s *S) TestServices(c *C) {
	storage, err := Conn()
	c.Assert(err, IsNil)
	services := storage.Services()
	collection := storage.Collection("services")
	c.Assert(services, DeepEquals, collection)
}

func (s *S) TestUsers(c *C) {
	storage, err := Conn()
	c.Assert(err, IsNil)
	users := storage.Users()
	collection := storage.Collection("users")
	c.Assert(users, DeepEquals, collection)
}

func (s *S) TestTeams(c *C) {
	storage, err := Conn()
	c.Assert(err, IsNil)
	teams := storage.Teams()
	collection := storage.Collection("teams")
	c.Assert(teams, DeepEquals, collection)
}
