package database

import (
	"database/sql/driver"
	entsql "github.com/facebook/ent/dialect/sql"
	"contrib.go.opencensus.io/integrations/ocsql"
	"github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"fmt"
	"context"
	"time"
)

var logger = getLogger()
var (
	TemporalHostPort ="localhost:7233"
	TemporalNamespace ="default"
)

func getLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Unable to initialize log at init()")
	}
	return logger
}

type connector struct {
	dsn string
}

func (c connector) Connect(context.Context) (driver.Conn, error) {
	return c.Driver().Open(c.dsn)
}

func (connector) Driver() driver.Driver {
	return ocsql.Wrap(
		mysql.MySQLDriver{},
		ocsql.WithAllTraceOptions(),
		ocsql.WithRowsClose(false),
		ocsql.WithRowsNext(false),
		ocsql.WithDisableErrSkip(true),
	)
}


func init() {
	configDir := os.Getenv("CONFIG_DIR")
    if configDir == "" {
        return // init of the test will read the config from their own init
    }
    viper.AddConfigPath(configDir)
    viper.SetConfigType("json")
    viper.SetConfigFile(configDir + "/config.json")
    viper.AutomaticEnv()
    TemporalHostPort = viper.GetString("worker_config.temporal_host_port")
    TemporalNamespace = viper.GetString("worker_config.namespace")
    if err := viper.ReadInConfig(); err != nil {
        logger.Panic("VIPER config read error", zap.Error(err))
    }
}

func GetDBConnection() (*entsql.Driver, error) {
	host := viper.GetString("host")
	port := viper.GetInt("port")
	username := viper.GetString("db_username")
    password := viper.GetString("db_password")
	databaseName := viper.GetString("database_name")
	dataSource := fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", username, password, host, port, databaseName)
	logger.Info("DB Connection", zap.String("host", host), zap.Int("port", port), zap.String("databaseName", databaseName))
	currDriver := Open(dataSource)
	return currDriver, nil
}

func Open(dsn string) *entsql.Driver  {
	db := sql.OpenDB(connector{dsn})
	maxOpenConnections := 25
    maxIdleConnections := 3
    maxConnectionLifeTime := 300*time.Second
    maxOpenConnections = viper.GetInt("max_open_connections")
    maxIdleConnections = viper.GetInt("max_idle_connections")
    maxConnectionLifeTime = time.Duration(viper.GetInt("max_connection_lifetime_seconds")) * time.Second
    db.SetMaxOpenConns(maxOpenConnections)
    db.SetMaxIdleConns(maxIdleConnections)
    db.SetConnMaxLifetime(maxConnectionLifeTime)
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB("mysql", db)
	return drv
}
