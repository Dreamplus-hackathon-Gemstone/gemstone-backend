package domain

import (
	"database/sql"
	"fmt"
	"gemstone-backend/internal/gemstone/config"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

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
		log.Panicf("Panic occured while opening postgresql client : %v\n", err)
	}

	//_ = client.Migrator().DropTable(&Genres{})
	//_ = client.Migrator().DropTable(&Proposal{})
	//_ = client.Migrator().DropTable(&User{})
	//_ = client.Migrator().DropTable(&Tokens{})
	//_ = client.Migrator().DropTable(&Proposal{})
	//_ = client.Migrator().DropTable(&MovieUser{})
	//_ = client.Migrator().DropTable(&ProposalUser{})

	err = client.Migrator().CreateTable(&Genres{})
	log.Printf("Create table Genres ... err : %v\n", err)

	err = client.Migrator().CreateTable(&Tokens{})
	log.Printf("Create table Tokens ... err : %v\n", err)

	err = client.Migrator().CreateTable(&User{})
	log.Printf("Create table User ... err : %v\n", err)

	err = client.Migrator().CreateTable(&Movie{})
	log.Printf("Create table Proposals ... err : %v\n", err)

	err = client.Migrator().CreateTable(&Proposal{})
	log.Printf("Create table Proposal ... err : %v\n", err)

	err = client.Migrator().CreateTable(&MovieUser{})
	log.Printf("Create table MovieUser ... err : %v\n", err)

	err = client.Migrator().CreateTable(&ProposalUser{})
	log.Printf("Create table ProposalMiners ... err : %v\n", err)

	return client
}
