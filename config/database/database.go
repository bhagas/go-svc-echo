package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type gormInstance struct {
	master *gorm.DB
}

// Master initialize DB for master data
func (g *gormInstance) Master() *gorm.DB {
	return g.master
}

// GormDatabase abstraction
type GormDatabase interface {
	Master() *gorm.DB
}

func InitGorm() GormDatabase {
	inst := new(gormInstance)

	newLogger := gormLogger.New(
		log.New(os.Stdout, "Query Statement \r\n", log.LstdFlags), // io writer
		gormLogger.Config{
			LogLevel:                  gormLogger.Info, // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,           // Don't include params in the SQL log
			Colorful:                  true,            // Disable color
		},
	)

	gormConfig := &gorm.Config{
		// enhance performance config
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 newLogger,
	}

	// username, password, host, port, database
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_MASTER_HOST"), os.Getenv("DB_MASTER_USERNAME"), os.Getenv("DB_MASTER_PASSWORD"),
		os.Getenv("DB_MASTER_NAME"), os.Getenv("DB_MASTER_PORT"))

	db, err := gorm.Open(postgres.Open(connection), gormConfig)
	if err != nil {
		// logger.E("cant connect to database")
		log.Printf("cant connect to database %v\n", err)
	}

	_, err = db.DB()
	if err != nil {
		// logger.E("Error setup pooling")
		log.Printf("Error setup pooling %v\n", err)
	}

	//sqldb.SetMaxIdleConns(5)
	//sqldb.SetMaxOpenConns(10)

	// if err := db.Use(otelgorm.NewPlugin()); err != nil {
	// 	panic(err)
	// }

	inst.master = db

	return inst
}
