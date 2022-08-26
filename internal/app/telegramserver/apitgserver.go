package telegramserver

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func NewDB(config *Config) (*sql.DB, error) {

	attempts := 5
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.NameDB)

	for ; attempts > 0; attempts-- {
		log.Println(fmt.Sprintf("Connecting to db. Attempts left %d", attempts))
		time.Sleep(500 * time.Millisecond)
		db, err := sql.Open("postgres", connStr)

		if err != nil {
			log.Println(fmt.Sprintf("Error: open db %s", err.Error()))
			continue
		}

		if err := db.Ping(); err != nil {
			log.Println(fmt.Sprintf("Error: ping db %s", err.Error()))
			continue
		}

		return db, nil
	}

	return nil, fmt.Errorf("Error: open db connection")
}
