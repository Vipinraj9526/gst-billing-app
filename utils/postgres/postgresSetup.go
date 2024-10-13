package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"gst-billing/commons/constants"
	"gst-billing/utils/configs"
	"log"
	"time"

	gormLogger "gorm.io/gorm/logger"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostGresConClient struct {
	GormDb *gorm.DB
	SqlDb  *sql.DB
}

var postgresClient *PostGresConClient

// InitPostgresDBConfig initializes the Postgres database configuration and establishes a connection to the database.
func InitPostgresDBConfig(ctx context.Context) error {
	PostgresConfig, err := configs.LoadConfig("configs/postgres.yml")
	if err != nil {
		return fmt.Errorf(constants.GetPostgresConfigError, err)
	}
	// Connection for postgres
	err = ConnectPostgresDatabase(ctx, PostgresConfig)
	if err != nil {
		return fmt.Errorf(constants.PostgresConnectionError, err)
	}
	return nil
}

// ConnectPostgresDatabase establishes a connection to the Postgres database using the provided configuration.
func ConnectPostgresDatabase(ctx context.Context, postgresConfig configs.Config) error {
	// Construct the DSN string
	dsn := fmt.Sprintf(constants.DNSString, postgresConfig.Postgres.Host, postgresConfig.Postgres.Port, postgresConfig.Postgres.Username, postgresConfig.Postgres.Password, postgresConfig.Postgres.Database, postgresConfig.Postgres.SSlMode, postgresConfig.Postgres.TimeZone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})
	if err != nil {

		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = sqlDB.PingContext(ctx)
	if err != nil {
		return err
	}
	log.Print(constants.PostgresConnectionSuccessful)

	SetPostgresClient(db, sqlDB)

	return nil
}

func SetPostgresClient(db *gorm.DB, sqlDB *sql.DB) {
	postgresClient = &PostGresConClient{GormDb: db, SqlDb: sqlDB}
}

func ClosePostgres(ctx context.Context) {
	if postgresClient != nil {
		err := postgresClient.SqlDb.Close()
		if err != nil {
			log.Print(constants.ClosePostgresClientError, zap.Error(err))
		}
	}
}

func GetPostGresClient() *PostGresConClient {
	return postgresClient
}
