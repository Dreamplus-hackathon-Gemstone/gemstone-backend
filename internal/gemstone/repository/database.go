package repository

import (
	"database/sql"
	"fmt"
	"gemstone-backend/internal/gemstone/config"
	"gemstone-backend/internal/gemstone/domain"
	"gemstone-backend/internal/gemstone/transactor"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Repository struct {
	ItemRepo     domain.IRepository
	TokenRepo    domain.IRepository
	MakerRepo    domain.IRepository
	MinerRepo    domain.IRepository
	CategoryRepo domain.IRepository
}

func NewRepository(db *gorm.DB) *Repository {
	item := NewProposalRepository(transactor.NewTransactor(db))
	miner := NewMinerRepository(transactor.NewTransactor(db))
	maker := NewMakerRepository(transactor.NewTransactor(db))
	token := NewTokenRepository(transactor.NewTransactor(db))
	category := NewGenreRepository(transactor.NewTransactor(db))
	return &Repository{ItemRepo: item, MinerRepo: miner, MakerRepo: maker, TokenRepo: token, CategoryRepo: category}
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

	_ = client.Migrator().DropTable(&domain.Genres{})
	_ = client.Migrator().DropTable(&domain.Proposals{})
	_ = client.Migrator().DropTable(&domain.Miners{})
	_ = client.Migrator().DropTable(&domain.Makers{})
	_ = client.Migrator().DropTable(&domain.Tokens{})
	_ = client.Migrator().DropTable(&domain.Movie{})
	_ = client.Migrator().DropTable(&domain.MovieMiners{})
	_ = client.Migrator().DropTable(&domain.ProposalMiners{})

	err = client.Migrator().CreateTable(&domain.Genres{})
	log.Printf("Create table Genres ... err : %v\n", err)
	err = client.Migrator().CreateTable(&domain.Tokens{})
	log.Printf("Create table Tokens ... err : %v\n", err)
	err = client.Migrator().CreateTable(&domain.Miners{})
	log.Printf("Create table Miners ... err : %v\n", err)
	err = client.Migrator().CreateTable(&domain.Makers{})
	log.Printf("Create table Makers ... err : %v\n", err)
	err = client.Migrator().CreateTable(&domain.Movie{})
	log.Printf("Create table Movie ... err : %v\n", err)
	err = client.Migrator().CreateTable(&domain.Proposals{})
	log.Printf("Create table Proposals ... err : %v\n", err)
	err = client.Migrator().CreateTable(&domain.MovieMiners{})
	log.Printf("Create table MovieMiners ... err : %v\n", err)
	err = client.Migrator().CreateTable(&domain.ProposalMiners{})
	log.Printf("Create table ProposalMiners ... err : %v\n", err)
	return client
}
