package easydb

type dbConfigParam func(*dbConfig)

func WithDataSourceName(DataSource string) dbConfigParam {
	return func(c *dbConfig) {
		c.DataSource = DataSource
	}
}
func WithUserName(UserName string) dbConfigParam {
	return func(c *dbConfig) {
		c.UserName = UserName
	}
}
func WithPassword(Password string) dbConfigParam {
	return func(c *dbConfig) {
		c.Password = Password
	}
}
func WithHost(Host string) dbConfigParam {
	return func(c *dbConfig) {
		c.Host = Host
	}
}
func WithPort(Port string) dbConfigParam {
	return func(c *dbConfig) {
		c.Port = Port
	}
}
func WithSchema(Schema string) dbConfigParam {
	return func(c *dbConfig) {
		c.Schema = Schema
	}
}
func NewMysqlConfig(params ...dbConfigParam) *dbConfig {
	config := &dbConfig{}
	for _, p := range params {
		p(config)
	}
	if config.UserName == "" {
		config.UserName = "root"
	}
	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.Port == "" {
		config.Port = "3306"
	}
	if config.DataSource == "" {
		config.DataSource = config.Host + config.Port + config.Schema
	}
	return config
}
func NewPgsqlConfig(params ...dbConfigParam) *dbConfig {
	config := &dbConfig{}
	for _, p := range params {
		p(config)
	}
	if config.UserName == "" {
		config.UserName = "postgres"
	}
	if config.Password == "" {
		config.Password = "postgres"
	}
	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.Port == "" {
		config.Port = "5432"
	}
	if config.DataSource == "" {
		config.DataSource = config.Host + config.Port + config.Schema
	}
	return config
}

type dbConfig struct {
	DataSource string
	UserName   string
	Password   string
	Host       string
	Port       string
	Schema     string
}
