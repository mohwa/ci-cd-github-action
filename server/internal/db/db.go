package db

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbArg = "charset=utf8mb4&parseTime=True&loc=Local"
)

var (
	db *gorm.DB
)

var (
	ErrorAlreadyInitialized  = errors.New("db is already initialized")
	ErrorEmptyDataSourceName = errors.New("data source name is empty")
	ErrorAlreadyClosed       = errors.New("db is already closed")
	ErrorNotInitialized      = errors.New("db is not initialized yet")
	ErrorDBInitFail          = errors.New("fail to initialize db")
	ErrorMigrationFail       = errors.New("fail to migrate table")
	ErrorNoTable             = errors.New("there is no table")
	ErrorTableOpenFail       = errors.New("fail to open table")
)

func isExistTable(name string) bool {
	if db == nil {
		return false
	}

	return db.Migrator().HasTable(name)
}

func Init(dbDataSourceName string) (err error) {
	if db != nil {
		return ErrorAlreadyInitialized
	}

	// for dev purpose, let it slide when dsn is empty
	if dbDataSourceName == "" {
		return nil
	}

	//dsn := "yanione:password@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = initDB(dbDataSourceName)

	if err != nil {
		return err
	}

	if err = migrateTables(); err != nil {
		return ErrorMigrationFail
	}

	return nil
}

func initDB(dsn string) (*gorm.DB, error) {
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

// Migrate the schema
func migrateTables() error {
	if db == nil {
		return ErrorNotInitialized
	}

	// 생성된 구조체를 통해, 관련 DB(테이블 등)가 생성된다.
	if err := db.AutoMigrate(
		Todo{},
		Settings{},
	); err != nil {
		return err
	}

	return nil
}

func Close() {
	if dbH, _ := db.DB(); dbH != nil {
		dbH.Close()
	}

	db = nil
}
