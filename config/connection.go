package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var sharedDB *sql.DB
var once sync.Once

func Connect() *sql.DB {

	defer func() {
		if sharedDB != nil {
			fmt.Println("\n\n-------------------")
			stats := sharedDB.Stats()
			fmt.Println("stats.OpenConnections : ", stats.OpenConnections)
			err := sharedDB.Ping()
			if err != nil {
				fmt.Println("ping err:", err)
			}
			fmt.Println("-------------------")
		}
	}()

	once.Do(func() {

		_, currentFile, _, _ := runtime.Caller(0)
		appDir := filepath.Dir(currentFile)

		envPath := filepath.Join(appDir, "../.env")
		if err := godotenv.Load(envPath); err != nil {
			log.Fatalf("Error loading .env file from %s: %v", envPath, err)
		}
		fmt.Println("=================ONCE.DO")
		MYSQL_USERNAME := os.Getenv("MYSQL_USERNAME")
		MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
		MYSQL_HOST := os.Getenv("MYSQL_HOST")
		MYSQL_PORT := os.Getenv("MYSQL_PORT")
		MYSQL_DB := os.Getenv("MYSQL_DB")
		db, err := sql.Open("mysql", MYSQL_USERNAME+":"+MYSQL_PASSWORD+"@tcp("+MYSQL_HOST+":"+MYSQL_PORT+")/"+MYSQL_DB+"?parseTime=false")

		defer func() {
			if r := recover(); r != nil {
				log.Printf("\n\n\nPANIC RECOVERED in Connect(): %v\n", r)
			}
		}()

		if err != nil {
			log.Println("Connect sql.Open failed:", err)
			panic(err)
		}

		db.SetConnMaxIdleTime(25)
		db.SetMaxIdleConns(10)
		db.SetConnMaxLifetime(5 * time.Minute)

		sharedDB = db

	})

	return sharedDB
}
