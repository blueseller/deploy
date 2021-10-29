package configure

import (
	"bytes"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

var config1_0 = `
version: 1.0
log:
  level: debug
db:
  mongo:
    addr: 10.0.0.1
    port: 3306
    username: lk
    password: xxxx
    parameters:
      k1: v1
      k2: v2
      k3: true
maptest:
  testk:
    tk1: tk1
    tk2: tk2
`

type MySuite struct{}

var _ = Suite(&MySuite{})

func (m *MySuite) TestParseConfig(c *C) {
	configVal, err := Parse(bytes.NewReader([]byte(config1_0)))
	// 是否正常解析
	c.Assert(err, IsNil)
	// 是否能够回去正确的值
	c.Assert(configVal.Db.Mongo.Addr, Equals, "10.0.0.1")

	// 环境变量优先级高于配置文件, 环境变量的值是否替换成功
	// 替换一个struct 里面的值, export DEPLOY_LOG_LOGLEVEL=info
	c.Assert(configVal.Log.LogLevel, Equals, LogLevel("info"))

	// 替换一个map里面的值, export DEPLOY_DB_MONGO_PARAMETERS_K1=vvvvv
	c.Assert(configVal.Db.Mongo.Parameters["k1"], Equals, "vvvvv")

	// 替换一个map里面的值, export DEPLOY_MAPTEST_TESTK_TK1=vvvvv
	c.Assert(configVal.MapTest["testk"]["tk1"], Equals, "vvvvv")

	c.Log(configVal.Db.Mongo.Addr, configVal.Db.Mongo.Parameters["k1"])
}
