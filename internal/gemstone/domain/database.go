package domain

import (
	"database/sql"
	"fmt"
	"gemstone-backend/internal/gemstone/config"
	"gemstone-backend/internal/gemstone/global"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Repository[R global.ReturnType, P global.Parameter] struct {
	ProposalRepo global.IRepository[P, R]
	TokenRepo    global.IRepository[P, R]
	MakerRepo    global.IRepository[P, R]
	MinerRepo    global.IRepository[P, R]
	GenreRepo    global.IRepository[P, R]
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
		log.Panicf("Panic occured while opening postgresql client : %v\n", err)
	}

	_ = client.Migrator().DropTable(&Genres{})
	_ = client.Migrator().DropTable(&Proposals{})
	_ = client.Migrator().DropTable(&Miners{})
	_ = client.Migrator().DropTable(&Makers{})
	_ = client.Migrator().DropTable(&Tokens{})
	_ = client.Migrator().DropTable(&Movie{})
	_ = client.Migrator().DropTable(&MovieMiners{})
	_ = client.Migrator().DropTable(&ProposalMiners{})

	err = client.Migrator().CreateTable(&Genres{})
	log.Printf("Create table Genres ... err : %v\n", err)
	err = client.Migrator().CreateTable(&Tokens{})
	log.Printf("Create table Tokens ... err : %v\n", err)
	err = client.Migrator().CreateTable(&Miners{})
	log.Printf("Create table Miners ... err : %v\n", err)
	err = client.Migrator().CreateTable(&Makers{})
	log.Printf("Create table Makers ... err : %v\n", err)
	err = client.Migrator().CreateTable(&Movie{})
	log.Printf("Create table Movie ... err : %v\n", err)
	err = client.Migrator().CreateTable(&Proposals{})
	log.Printf("Create table Proposals ... err : %v\n", err)
	err = client.Migrator().CreateTable(&MovieMiners{})
	log.Printf("Create table MovieMiners ... err : %v\n", err)
	err = client.Migrator().CreateTable(&ProposalMiners{})
	log.Printf("Create table ProposalMiners ... err : %v\n", err)
	return client
}
