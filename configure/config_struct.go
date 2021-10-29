package configure

type LogLevel string
type Version string

type Configuration struct {
	Version Version `yaml: "version,omitempty"`
	Log     struct {
		LogLevel LogLevel `yaml: "level,omitempty"`

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

	//MapTest map[string]map[string]interface{} `yaml: "maptest", omitempty`
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
