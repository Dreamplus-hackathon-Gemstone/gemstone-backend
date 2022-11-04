package context

import (
	"database/sql"
	"fmt"
	"gemstone-backend/internal/api/user"
	"gemstone-backend/internal/app/database"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// AppCtx ...
type AppCtx interface {
	DB() *gorm.DB
}

type appCtxData struct {
	db         *gorm.DB
	transactor database.ITransactor
}

func (ctx *appCtxData) DB() *gorm.DB {
	return ctx.db
}

// Initialize ...
func Initialize(config Config) (AppCtx, error) {
	db, err := createDb(config)
	if err != nil {
		return nil, err
	}
	return createAppCtx(db)
}

func createAppCtx(db *gorm.DB) (AppCtx, error) {
	transactor := database.NewTransactor(db)
	return &appCtxData{db: db, transactor: transactor}, nil
}

func createDb(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul", config.Host, config.User, config.Password, config.Name, config.Port)
	log.Printf("Create Database for %s\n", dsn)
	sqldb, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Panicf("Open database... panic!: %v", err)
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
		},
	)

	client, err := gorm.Open(postgres.New(postgres.Config{
		Conn:                 sqldb,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: newLogger,
	})

	createTable(client.Migrator())
	log.Println()
	return client, err
}

func createTable(m gorm.Migrator) {
	miner := &user.Miner{}
	maker := &user.Maker{}

	if err := m.CreateTable(miner); err != nil {
		log.Fatalf("CreateTable failed in %s : %v", "miner", err)
	}
	if err := m.CreateTable(maker); err != nil {
		log.Fatalf("CreateTable failed in %s : %v", "maker", err)
	}
}
