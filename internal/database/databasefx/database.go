package databasefx

import (
	"context"
	"fmt"

	raas "github.com/raas-app/stocks"
	"github.com/raas-app/stocks/internal/database"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func provideDatabase(lc fx.Lifecycle, config *raas.Config, logger *zap.Logger) (database.Database, error) {
	termCtx := context.Background()
	mainSqlx, closeDB, err := database.ConnectLoop(termCtx, config.Database, logger)
	if err != nil {
		return database.Database{}, err
	}
	lc.Append(fx.StopHook(closeDB))

	// replicaSqlx, closeReplicaDB, err := database.ConnectLoop(termCtx, config.Database, "replica", logger)
	// if err != nil {
	// 	return database.Database{}, err
	// }
	// lc.Append(fx.StopHook(closeReplicaDB))

	rw, err := database.NewCommonDB(mainSqlx, "mysql_main", logger.Named("mysql_main_common"))
	if err != nil {
		return database.Database{}, fmt.Errorf("err newCommonDB, %w", err)
	}
	// ro, err := database.NewCommonDB(replicaSqlx, "mysql_replica", logger.Named("mysql_replica_common"))
	// if err != nil {
	// 	return database.Database{}, fmt.Errorf("err newCommonDB, %w", err)
	// }

	return database.Database{
		RW:     rw,
		Pinger: mainSqlx,
	}, nil
}

var Module = fx.Module("database",
	fx.Provide(provideDatabase),
	fx.Invoke(func(_ database.Database) {}),
)
