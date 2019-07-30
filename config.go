package easydb

type dbConfigParam func(*DBConfig)

// WithDataSourceName set data source name option
func WithDataSourceName(DataSource string) dbConfigParam {
	return func(c *DBConfig) {
		c.DataSource = DataSource
	}
}

// WithUserName set username option
func WithUserName(UserName string) dbConfigParam {
	return func(c *DBConfig) {
		c.UserName = UserName
	}
}

// WithPassword set password option
func WithPassword(Password string) dbConfigParam {
	return func(c *DBConfig) {
		c.Password = Password
	}
}

// WithHost set host option
func WithHost(Host string) dbConfigParam {
	return func(c *DBConfig) {
		c.Host = Host
	}
}

// WithPort set port option
func WithPort(Port string) dbConfigParam {
	return func(c *DBConfig) {
		c.Port = Port
	}
}

// WithSchema set schema option
func WithSchema(Schema string) dbConfigParam {
	return func(c *DBConfig) {
		c.Schema = Schema
	}
}

// WithMaxOpenConns set schema option
func WithMaxOpenConns(MaxOpenConns int) dbConfigParam {
	return func(c *DBConfig) {
		c.MaxOpenConns = MaxOpenConns
	}
}

// WithMaxIdleConns set schema option
func WithMaxIdleConns(MaxIdleConns int) dbConfigParam {
	return func(c *DBConfig) {
		c.MaxIdleConns = MaxIdleConns
	}
}

// WithConnMaxLifetime set schema option
func WithConnMaxLifetime(ConnMaxLifetime int) dbConfigParam {
	return func(c *DBConfig) {
		c.ConnMaxLifetime = ConnMaxLifetime
	}
}

// WithConnParams set schema option
func WithConnParams(ConnParams string) dbConfigParam {
	return func(c *DBConfig) {
		c.ConnParams = ConnParams
	}
}

// WithEnableMigrate set schema option
func WithEnableMigrate(EnableMigrate bool) dbConfigParam {
	return func(c *DBConfig) {
		c.EnableMigrate = EnableMigrate
	}
}

// WithMigrateDir set schema option
func WithMigrateDir(MigrateDir string) dbConfigParam {
	return func(c *DBConfig) {
		c.MigrateDir = MigrateDir
	}
}

// NewMysqlConfig create new mysql config
func NewMysqlConfig(params ...dbConfigParam) *DBConfig {
	config := &DBConfig{}
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
	if config.MigrateDir == "" {
		config.MigrateDir = "./migrations"
	}
	return config
}

// NewPgsqlConfig create pgsql config
func NewPgsqlConfig(params ...dbConfigParam) *DBConfig {
	config := &DBConfig{}
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
	if config.MigrateDir == "" {
		config.MigrateDir = "./migrations"
	}
	return config
}

// DBConfig db config struct
type DBConfig struct {
	DataSource      string
	UserName        string
	Password        string
	Host            string
	Port            string
	Schema          string
	ConnParams      string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int
	EnableMigrate   bool
	MigrateDir      string
}
