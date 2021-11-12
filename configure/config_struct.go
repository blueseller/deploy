package configure

type LogLevel string
type Version string

type Configuration struct {
	Version Version `yaml: "version,omitempty"`
	Log     struct {
		LogType string `yaml: "log_type,omitempty"`

		LogLevel LogLevel `yaml: "loglevel,omitempty"`

		// 支持 text、json、logstash
		Formatter string `yaml: "formatter,omitempty"`

		Fields map[interface{}]interface{} `yaml: "fields,omitempty"`

		Hooks LogHook `yaml: "hooks,omitempty"`
	} `yaml: "log,omitempty"`

	// 设定日志打印级别
	LogLevel LogLevel `yaml: "level,omitempty"`

	Db struct {
		Mysql DbStruct `yaml: "mysql", omitempty`
		PG    DbStruct `yaml: "pg", omitempty`
		Mongo DbStruct `yaml: "mongo, omitempty"`
	} `yaml: "db", omitempty`

	// 测试用key， 后期删除
	MapTest map[string]map[string]interface{} `yaml: "maptest", omitempty`

	CmdFlow CmdFlow `yaml: "cmdflow", omitempty`
}

type CmdStep int

type CmdFlow map[CmdStep]FlowCommand

type FlowCommand struct {
	Num          CmdStep   `yaml: "num", omitempty`
	Name         string    `yaml: "name", omitempty`
	Desc         string    `yaml: "desc", omitempty`
	NextCmdSteps []CmdStep `yaml: "nextcmdsteps", omitempty`
	Commands     []Command `yaml: "commands", omitempty`
}

type Command struct {
	Name   string   `yaml: "name", omitempty`
	Desc   string   `yaml: "desc", omitempty`
	Hander string   `yaml: "hander", omitempty`
	Args   []string `yaml: "args", omitempty`
}

type DbStruct struct {
	Addr       string                 `yaml: "addr", omitempty`
	Port       string                 `yaml: "port", omitempty`
	UserName   string                 `yaml: "username", omitempty`
	Password   string                 `yaml: "password", omitempty`
	Parameters map[string]interface{} `yaml: "parameters", omitempty`
}

// log hook 可以在打印某些特定的日志情况下, 发送邮件, 或者做某些逻辑处理
type LogHook struct {
	Disabled bool `yaml: "disabled,omitempty"`

	Type string `yaml: "yaml:type,omitempty"`

	Levels []string `yaml: "levels, omitempty"`

	MailOptions MailOptions `yaml:"options,omitempty"`
}

type MailOptions struct {
	SMTP struct {
		Addr string `yaml: addr,omitempty`

		UserName string `yaml: username,omitempty`

		Password string `yaml: passwod,omitempty`

		// 不做认证，跳过登录
		Insecure bool `yaml: insecure,omitempty`
	} `yaml: smtp,omitempty`

	From string `yaml: from,omitempty`

	To []string `yaml: to, omitempty`
}

type v0_1Configuration Configuration
