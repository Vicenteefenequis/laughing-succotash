package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"laughing-succostash/internal/core/domain"
	"log"
	"os"
	"strconv"
	"strings"
)

type Options = []interface{}

type Database struct {
	User    string
	Pass    string
	Host    string
	Port    int
	DbName  string
	Options *Options
	Db      *gorm.DB
}

func NewDatabase() *Database {
	port, _ := strconv.Atoi(os.Getenv("MYSQL_PORT"))

	return &Database{
		User:    os.Getenv("MYSQL_USER"),
		Pass:    os.Getenv("MYSQL_PASSWORD"),
		Host:    os.Getenv("MYSQL_HOST"),
		Port:    port,
		DbName:  os.Getenv("MYSQL_DATABASE"),
		Options: &[]interface{}{},
	}
}

func (db *Database) Connect() *gorm.DB {

	connection, err := gorm.Open(
		mysql.Open(db.MountDsn()),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatalf("Ocorreu um erro ao conectar a base de dados")
	}

	db.Db = connection
	db.Db.AutoMigrate(&domain.User{})
	db.Db.AutoMigrate(&domain.Category{})

	return connection
}

func (db *Database) MountDsn() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		db.User,
		db.Pass,
		db.Host,
		db.Port,
		db.DbName,
		db.mountOptions(),
	)
}

func (db *Database) mountOptions() string {
	options := ""

	for option, value := range *db.Options {
		options += fmt.Sprintf("%v=%v&", option, value)
	}

	return strings.TrimSuffix(options, "&")
}
