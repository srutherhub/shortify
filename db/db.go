package db

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/teris-io/shortid"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite", "shortify.db")
	if err != nil {
		panic(err)
	}

	_, err = DB.Exec("PRAGMA journal_mode=WAL;")
	if err != nil {
		log.Fatal("Failed to enable WAL mode:", err)
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS urls (
		id TEXT PRIMARY KEY,
		url TEXT NOT NULL,
		created_at NUMBER NOT NULL,
		click_count INTEGER NOT NULL DEFAULT 0
	)`)

	if err != nil {
		panic(err)
	}

}

func GetURLFromID(id string) (string, error) {
	var url string
	err := DB.QueryRow(`SELECT url from urls WHERE id = ?`, id).Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}

func InsertNewURL(url string) (string, error) {
	id := CreateUniqueID()
	now := time.Now().Unix()

	if !isURLHttps(url) {
		return "", errors.New("URL is not Https")
	}

	_, err := DB.Exec(`INSERT into urls(id,url,created_at) VALUES(?,?,?)`, id, url, now)

	if err != nil {
		return "", err
	}
	return id, nil
}

func IncrementURLClickCount(id string) error {
	_, err := DB.Exec(`UPDATE urls SET click_count = click_count + 1 WHERE id = ?`, id)
	return err
}

func CreateUniqueID() string {
	id := shortid.MustGenerate()
	return id
}

func isURLHttps(url string) bool {
	firstEightChars := url[:8]
	if firstEightChars == "https://" {
		return true
	} else {
		return false
	}
}
