package database

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
)

func PostgresInit() *gorm.DB {
	var DB *gorm.DB
	var err error
	var once sync.Once
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             0,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
		},
	)
	once.Do(func() {
		DB, err = gorm.Open(postgres.New(postgres.Config{
			DriverName: "postgres",
			DSN:        os.Getenv("POSTGRES_URL"),
		}), &gorm.Config{Logger: newLogger})
		if err != nil {
			log.Fatalf("Unable to connect to database: %v\n", err)
		}
	})

	log.Println("Database connection established.")

	return DB
}
