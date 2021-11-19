package configure

var defaultConfig *Configuration

func init() {
	defaultConfig = new(Configuration)
}

func SetConfig(cfg *Configuration) {
	defaultConfig = cfg
}

func GetConfig() *Configuration {
	return defaultConfig
}
