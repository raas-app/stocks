package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	raas "github.com/raas-app/stocks"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const (
	defaultTimeout       = 5 * time.Second
	driverName           = "mysql"
	dateTimeFormatInJSON = time.RFC3339
	maxOpenConns         = 200
	maxIdleConns         = 200
	connMaxLifetime      = 3 * time.Minute
	connMaxIdleTime      = 4 * time.Minute
)

type Database struct {
	RW     Common
	Pinger DBPinger
}

// ConnectLoop takes config and specified database credentials as input, returning *sqlx.DB handle for interactions
// with database
// It tries to connect in a loop because service can be healthy and up before database is ready in docker compose for example.
func ConnectLoop(
	ctx context.Context,
	databaseConfig raas.DatabaseConfig,
	logger *zap.Logger) (db *sqlx.DB, closeFunc func() error, err error) {
	if logger == nil {
		return nil, nil, errors.New("database: provided logger is nil")
	}

	dbCfg, err := getDBConfig(databaseConfig)
	if err != nil {
		return nil, nil, err
	}
	cfg, err := newMysqlCfg(dbCfg)
	if err != nil {
		return nil, nil, err
	}

	logger = logger.With(
		zap.String("driver", driverName),
		zap.String("addr", cfg.Addr),
		zap.String("database", cfg.DBName),
		zap.String("username", cfg.User),
	)
	logger.Info("connecting to database")
	dsn := cfg.FormatDSN()
	db, err = establishConnection(ctx, driverName, dsn, dbCfg)
	if err == nil {
		return db, db.Close, nil
	}

	logger.Error("failed to connect to the database", zap.Error(err))

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	timeoutExceeded := time.After(dbCfg.Timeout)

	for {
		select {
		case <-timeoutExceeded:
			return nil, nil, fmt.Errorf("db connection failed after %f seconds timeout", dbCfg.Timeout.Seconds())
		case <-ticker.C:
			db, err := establishConnection(ctx, driverName, dsn, dbCfg)
			if err == nil {
				return db, db.Close, nil
			}
			logger.Error("failed to connect to the database", zap.Error(err))

		case <-ctx.Done():
			return nil, nil, ctx.Err()
		}
	}
}

func getDBConfig(databaseConfig raas.DatabaseConfig) (*raas.DatabaseConfig, error) {
	dbCfg := databaseConfig

	if dbCfg.Timeout == 0 {
		dbCfg.Timeout = defaultTimeout
	}

	if dbCfg.ConnectionsMaxLifetime == 0 {
		dbCfg.ConnectionsMaxLifetime = connMaxLifetime
	}
	if dbCfg.ConnectionMaxIdleTime == 0 {
		dbCfg.ConnectionMaxIdleTime = connMaxIdleTime
	}
	if dbCfg.MaxOpenConnections == 0 {
		dbCfg.MaxOpenConnections = maxOpenConns
	}
	if dbCfg.MaxIdleConnections == 0 {
		dbCfg.MaxIdleConnections = maxIdleConns
	}

	return &dbCfg, nil
}

func newMysqlCfg(dbSetting *raas.DatabaseConfig) (*mysql.Config, error) {
	cfg := mysql.NewConfig()
	cfg.Params = make(map[string]string)
	cfg.Net = dbSetting.Net
	cfg.Addr = dbSetting.Host
	cfg.User = dbSetting.User
	cfg.Passwd = dbSetting.Password
	cfg.DBName = dbSetting.Name
	cfg.ParseTime = true
	cfg.MultiStatements = true
	cfg.Timeout = dbSetting.Timeout
	cfg.RejectReadOnly = true

	return cfg, nil
}

func establishConnection(ctx context.Context, driverName string, dsn string, dbCfg *raas.DatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, driverName, dsn) // sqlx.Connect performs ping under the hood
	if err != nil {
		return nil, fmt.Errorf("database: problem while trying to connect to the database, %w", err)
	}

	db.SetMaxOpenConns(dbCfg.MaxOpenConnections)
	db.SetMaxIdleConns(dbCfg.MaxIdleConnections)
	db.SetConnMaxLifetime(dbCfg.ConnectionsMaxLifetime)
	db.SetConnMaxIdleTime(dbCfg.ConnectionMaxIdleTime)

	return db, nil
}
