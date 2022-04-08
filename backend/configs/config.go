package configs

import "github.com/MatheusHenrique129/application-in-go/util"

// Constants

const (
	DBDefaultHost  = "tcp(localhost:3306)"
	DBDefaultUser  = "root"
	DBDefaultPass  = "Sua Senha do banco de dados!!"
	DBDefaultName  = "bdGolang"
	DBMaxOpenConns = 100
	DBMaxIdleConns = 50
)

// Config This struct holds the configuration parameters.
type Config struct {
	dbHost string `envconfig:"HOST"`
	dbUser string `envconfig:"USER"`
	dbPass string `envconfig:"PASSWORD"`
	dbName string `envconfig:"DB_NAME"`

	logger *util.Logger
}

// GetDatabaseHost Returns the database hostname.
func (c *Config) GetDatabaseHost() string {
	return c.dbHost
}

// GetDatabaseUser Returns the database user.
func (c *Config) GetDatabaseUser() string {
	return c.dbUser
}

// GetDatabasePass Returns the database password.
func (c *Config) GetDatabasePass() string {
	return c.dbPass
}

// GetDatabaseName Returns the database name.
func (c *Config) GetDatabaseName() string {
	return c.dbName
}

//Loads the configuration.
func (c *Config) LoadConf() {
	c.dbHost = DBDefaultHost
	c.dbUser = DBDefaultUser
	c.dbPass = DBDefaultPass
	c.dbName = DBDefaultName
}

func NewConfig() *Config {
	c := &Config{
		logger: util.NewLogger("Config"),
	}

	c.LoadConf()

	return c
}
