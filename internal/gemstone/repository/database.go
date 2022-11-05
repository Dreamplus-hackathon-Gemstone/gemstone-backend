package repository

import (
	"database/sql"
	"fmt"
	"gemstone-backend/internal/gemstone/config"
	"gemstone-backend/internal/gemstone/domain"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type ITransactor interface {
	GetTx() *gorm.DB
	Execute() error
	ExecuteContext() error
}

type Transactor struct {
	Client     *gorm.DB
	Repository Repository
}

type Repository struct {
	ItemRepo     domain.Repository
	TokenRepo    domain.Repository
	MakerRepo    domain.Repository
	MinerRepo    domain.Repository
	CategoryRepo domain.Repository
}

func NewRepository(db *gorm.DB) *Repository {
	item := NewItemRepository(db)
	miner := NewMinerRepository(db)
	maker := NewMakerRepository(db)
	token := NewTokenRepository(db)
	category := NewCategoryRepository(db)
	return &Repository{ItemRepo: item, MinerRepo: miner, MakerRepo: maker, TokenRepo: token, CategoryRepo: category}
}

func NewTransactor(client *gorm.DB) *Transactor {
	return &Transactor{Client: client}
}

func (t Transactor) GetTx() *gorm.DB {
	//TODO implement me
	panic("implement me")
}

func (t Transactor) Execute() error {
	//TODO implement me
	panic("implement me")
}

func (t Transactor) ExecuteContext() error {
	//TODO implement me
	panic("implement me")
}

func NewClient(config *config.Config) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
		},
	)

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		config.USER, config.PASSWORD, config.NAME, config.PORT)
	database, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Panicf("Panic occured while opening database : %v\n", err)
	}
	client, err := gorm.Open(postgres.New(postgres.Config{
		Conn:                 database,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Panicf("Panic occured while opening postgresql client\n")
	}
	_ = client.Migrator().CreateTable(&domain.Genre{})
	_ = client.Migrator().CreateTable(&domain.Proposal{})
	_ = client.Migrator().CreateTable(&domain.Miner{})
	_ = client.Migrator().CreateTable(&domain.Maker{})
	_ = client.Migrator().CreateTable(&domain.Token{})
	return client
}
