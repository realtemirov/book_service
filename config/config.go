package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	DebugMode   = "debug"
	TestMode    = "test"
	ReleaseMode = "release"
)

type Config struct {
	ServiceName string
	Environment string
	Version     string

	HTTPPort   int
	HTTPSchema string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	PostgresSSLMode  string

	DefaultOffset int
	DefaultLimit  int

	PasscodePool   string
	PasscodeLength int

	ServiceHost string
	GRPCPort    int

	PostgresMaxConnections  int
	PostgresConnMaxIdleTime int
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found: " + err.Error())
	}

	config := Config{}
	config.ServiceName = cast.ToString(getOrDefault("SERVICE_NAME", "go-service"))
	config.Environment = cast.ToString(getOrDefault("ENVIRONMENT", "development"))
	config.Version = cast.ToString(getOrDefault("VERSION", "0.0.1"))

	config.HTTPPort = cast.ToInt(getOrDefault("HTTP_PORT", 8080))
	config.HTTPSchema = cast.ToString(getOrDefault("HTTP_SCHEMA", "http"))

	config.PostgresHost = cast.ToString(getOrDefault("POSTGRES_HOST", "0.0.0.0"))
	config.PostgresPort = cast.ToInt(getOrDefault("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrDefault("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrDefault("POSTGRES_PASSWORD", "postgres"))
	config.PostgresDatabase = cast.ToString(getOrDefault("POSTGRES_DATABASE", config.ServiceName))
	config.PostgresSSLMode = cast.ToString(getOrDefault("POSTGRES_SSLMODE", "disable"))
	config.DefaultOffset = cast.ToInt(getOrDefault("DEFAULT_OFFSET", 0))
	config.DefaultLimit = cast.ToInt(getOrDefault("DEFAULT_LIMIT", 10))

	config.PasscodePool = cast.ToString(getOrDefault("PASSCODE_POOL", "0123456789"))
	config.PasscodeLength = cast.ToInt(getOrDefault("PASSCODE_LENGTH", 6))

	config.ServiceHost = cast.ToString(getOrDefault("BOOK_SERVICE_HOST", "localhost"))
	config.GRPCPort = cast.ToInt(getOrDefault("BOOK_GRPC_PORT", 3001))

	config.PostgresMaxConnections = cast.ToInt(getOrDefault("POSTGRES_MAX_CONNECTIONS", 5))
	config.PostgresConnMaxIdleTime = cast.ToInt(getOrDefault("POSTGRES_CONN_MAX_IDLE_TIME", 10))

	return config

}

func getOrDefault(key string, defaultValue interface{}) interface{} {
	val, ok := os.LookupEnv(key)

	if !ok {
		return defaultValue
	}

	return val
}
