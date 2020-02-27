package postgres_config

import (
	"go.uber.org/zap"
	"os"
)

const (
	envPgHost = "PG_HOST"
	envPgPort = "PG_PORT"
	envPgUser = "PG_USER"
	envPgPass = "PG_PASS"
	envPgDB   = "PG_DBNAME"
	envPgSSL  = "PG_SSL"

	defaultPgHost = "localhost"
	defaultPgPort = "5432"
	defaultPgUser = "postgres_config"
	defaultPgPass = "Dbrnjh777"
	defaultPgDB   = "postgres_config"
	defaultPgSSL  = "disable"
)

type Config struct {
	PgHost, PgPort, PgUser, PgPass, PgDB, PgSSL string
}

func NewConfig(logger *zap.Logger) *Config {
	var (
		host, port, user, pass, dbname, ssl string
	)

	if host = os.Getenv(envPgHost); host == "" {
		logger.Info(
			"environment variable is not set - default is set",
			zap.String(envPgHost, "If not installed, there may be an error when connecting to the DB"),
		)
		host = defaultPgHost
	}

	if port = os.Getenv(envPgPort); port == "" {
		logger.Info(
			"environment variable is not set - default is set",
			zap.String(envPgPort, "If not installed, there may be an error when connecting to the DB"),
		)
		port = defaultPgPort
	}

	if user = os.Getenv(envPgUser); user == "" {
		logger.Info(
			"environment variable is not set - default is set",
			zap.String(envPgUser, "If not installed, there may be an error when connecting to the DB"),
		)
		user = defaultPgUser
	}

	if host = os.Getenv(envPgHost); host == "" {
		logger.Info(
			"environment variable is not set - default is set",
			zap.String(envPgHost, "If not installed, there may be an error when connecting to the DB"),
		)
		host = defaultPgHost
	}

	if pass = os.Getenv(envPgPass); pass == "" {
		logger.Info(
			"environment variable is not set - default is set",
			zap.String(envPgPass, "If not installed, there may be an error when connecting to the DB"),
		)
		pass = defaultPgPass
	}

	if dbname = os.Getenv(envPgDB); dbname == "" {
		logger.Info(
			"environment variable is not set - default is set",
			zap.String(envPgDB, "If not installed, there may be an error when connecting to the DB"),
		)
		dbname = defaultPgDB
	}

	if ssl = os.Getenv(envPgSSL); ssl == "" {
		logger.Info(
			"environment variable is not set - default is set",
			zap.String(envPgSSL, "If not installed, there may be an error when connecting to the DB"),
		)
		ssl = defaultPgSSL
	}

	return &Config{
		PgHost: host, PgPort: port, PgUser: user, PgPass: pass, PgDB: dbname, PgSSL: ssl,
	}
}

func (c *Config) String() string {
	return "host=" + c.PgHost + " port=" + c.PgPort +
		" user=" + c.PgUser + " password=" + c.PgPass +
		" dbname=" + c.PgDB + " sslmode=" + c.PgSSL
}
