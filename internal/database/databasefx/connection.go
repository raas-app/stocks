package databasefx

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //nolint
	raas "github.com/raas-app/stocks"
	"go.uber.org/zap"
)

type Connection interface {
	Connect() error
}

type databaseConnection struct {
	db     *sql.DB
	logger *zap.Logger
	config *raas.Config
}

func NewConnectionBuilder(logger *zap.Logger, config *raas.Config) Connection {
	var db *sql.DB

	return &databaseConnection{
		db:     db,
		logger: logger,
		config: config,
	}
}

func InvokeDatabaseConnection(c Connection) error {
	err := c.Connect()
	if err != nil {
		return err
	}
	return nil
}

func (c *databaseConnection) Connect() error {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?timeout=%s",
		c.config.Database.User,
		c.config.Database.Password,
		c.config.Database.Net,
		c.config.Database.Host,
		c.config.Database.Port,
		c.config.Database.Name,
		c.config.Database.Timeout,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	c.db = db

	return nil
}
