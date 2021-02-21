package storage

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"time"
)

// DbStore ...
type DbStore struct {
	*sqlx.DB
}

// Opening a storage and save the reference to `Database` struct.
func InitializeDB() *DbStore {
	//dataSourceName := fmt.Sprintf(core.Database.Username + ":" + core.Database.Password + "@/" + core.Database.Database)
	cnf := fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_DATABASE"))

	var err error
	var db *sqlx.DB

	if db, err = sqlx.Open("mysql", cnf); err != nil {
		log.Fatal(err)
	}

	retryCount := 30
	for {
		err := db.Ping()
		if err != nil {
			if retryCount == 0 {
				log.Fatalf("Not able to establish connection to database")
			}

			log.Printf(fmt.Sprintf("Could not connect to database. Wait 2 seconds. %d retries left...", retryCount))
			retryCount--
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}

	if errPing := db.Ping(); errPing != nil {
		log.Fatal(errPing)
	}
	return &DbStore{
		db,
	}
}
