package configure

import (
	"bytes"
	"os"
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
cmdflow:
  1:
    name: "start cmd"
    desc: "start cmd"
    next_cmd_steps:
      - 2
      - 3
  2:
    name: "deploy"
    desc: "整体部署"
    commands:
      - name: "deploy all"
        desc: "部署全部服务"
      - name: "deploy service"
        desc: "部署一个服务,包含服务和migration"
      - name: "deploy service only"
        desc: "只部署服务"
      - name: "deploy migration only"
        desc: "只部署服务的migration"
  3:
    name: "config"
    desc: "配置"
    commands:
      - name: "deploy config all"
        desc: "刷新配置"
      - name: "deploy config only"
        desc: "刷新一个服务"
      - name: "deploy config diff"
        desc: "查看配置差异"
      - name: "deploy config diff only"
        desc: "查看某一个配置差异"
`

type MySuite struct{}

var _ = Suite(&MySuite{})

func ENV_INIT() {
	// 设定一些env
	os.Setenv("DEPLOY_DB_MONGO_PARAMETERS_K1", "vvvvv")
	os.Setenv("DEPLOY_MAPTEST_TESTK_TK1", "vvvvv")
	os.Setenv("DEPLOY_LOG_LOGLEVEL", "info")
}

func ENV_DELETE() {
	// 删除一些env
	os.Unsetenv("DEPLOY_DB_MONGO_PARAMETERS_K1")
	os.Unsetenv("DEPLOY_MAPTEST_TESTK_TK1")
	os.Unsetenv("DEPLOY_LOG_LOGLEVEL")
}

func (m *MySuite) SetUpSuite(c *C) {
	ENV_INIT()
}

func (m *MySuite) TearDownSuite(c *C) {
	ENV_DELETE()
}

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

	//log.Info(configVal)
}
