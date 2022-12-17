package config

import (
	"github.com/MatheusHenrique129/application-in-go/internal/util"
)

// Constants
const (
	// Database
	DriverName     = "mysql"
	DBDefaultHost  = "tcp(localhost:3306)"
	DBDefaultUser  = "root"
	DBDefaultPass  = "Math@2109"
	DBDefaultName  = "bdgolang"
	DBMaxOpenConns = 100
	DBMaxIdleConns = 50
)

// Env This struct is required so we can parse the data from the Environment. We have them in another struct
// so we can hide them from Config.
type Env struct {
	// Port specifies the port this application will be listening at
	Port string `envconfig:"PORT" required:"true" default:"8080"`
}

// Config This struct holds the configuration parameters.
type Config struct {
	port                         string
	appShutdownGraceMilliseconds int

	// Database credentials
	databaseHost string
	databaseUser string
	databasePass string
	databaseName string
	logger       *util.Logger
}

// GetPort Returns the application port.
func (c *Config) GetPort() string {
	return c.port
}

// GetAppShutdownGraceMilliseconds Determines how many seconds server has to finish when shutting down
func (c *Config) GetAppShutdownGraceMilliseconds() int {
	return c.appShutdownGraceMilliseconds
}

// GetDatabaseHost Returns the database hostname.
func (c *Config) GetDatabaseHost() string {
	return c.databaseHost
}

// GetDatabaseUser Returns the database user.
func (c *Config) GetDatabaseUser() string {
	return c.databaseUser
}

// GetDatabasePass Returns the database password.
func (c *Config) GetDatabasePass() string {
	return c.databasePass
}

// GetDatabaseName Returns the database name.
func (c *Config) GetDatabaseName() string {
	return c.databaseName
}

// Load the configuration.
func (c *Config) Load() {
	confEnv := &Env{}

	// Load config parameters
	port := confEnv.Port
	if port == "" {
		port = "8080"
	}

	// Load config parameters from env
	c.port = port
	c.databaseHost = DBDefaultHost
	c.databaseUser = DBDefaultUser
	c.databasePass = DBDefaultPass
	c.databaseName = DBDefaultName
}

func NewConfig() *Config {
	c := &Config{
		logger: util.NewLogger("Config"),
	}

	c.Load()
	return c
}
