package tidblite


import (
	"context"
	"fmt"
	"testing"
	"time"

	tidblite "github.com/WangXiangUSTC/tidb-lite"
	. "github.com/pingcap/check"
)

func TestClient(t *testing.T) {
	TestingT(t)
}

var _ = Suite(&testExampleSuite{})

type testExampleSuite struct{}

func (t *testExampleSuite) TestGetRowCount(c *C) {
	tidbServer, err := tidblite.NewTiDBServer(tidblite.NewOptions(c.MkDir()))
	c.Assert(err, IsNil)

	dbConn, err := tidbServer.CreateConn()
	fmt.Printf("dbConn: %+v, err: %+v", dbConn, err)
	c.Assert(err, IsNil)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = dbConn.ExecContext(ctx, "create database example_test")
	c.Assert(err, IsNil)
	_, err = dbConn.ExecContext(ctx, "create table example_test.t(id int primary key, name varchar(24))")
	c.Assert(err, IsNil)
	_, err = dbConn.ExecContext(ctx, "insert into example_test.t values(1, 'a'),(2, 'b'),(3, 'c')")
	c.Assert(err, IsNil)

	count, err := GetRowCount(ctx, dbConn, "example_test", "t", "id > 2")
	c.Assert(err, IsNil)
	c.Assert(count, Equals, int64(1))

	count, err = GetRowCount(ctx, dbConn, "example_test", "t", "")
	c.Assert(err, IsNil)
	c.Assert(count, Equals, int64(3))
	tidbServer.Close()

	tidbServer2, err := tidblite.NewTiDBServer(tidblite.NewOptions(c.MkDir()))
	c.Assert(err, IsNil)
	defer tidbServer2.Close()

	dbConn2, err := tidbServer2.CreateConn()
	c.Assert(err, IsNil)
	_, err = dbConn2.ExecContext(ctx, "create database example_test")
	c.Assert(err, IsNil)
}