package db

import (
	"database/sql"
	"log"

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
		url TEXT NOT NULL
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

func InsertNewURL(url string) error {
	id := CreateUniqueID()
	_, err := DB.Exec(`INSERT into urls(id,url) VALUES(?,?)`, id, url)
	return err
}

func CreateUniqueID() string {
	id := shortid.MustGenerate()
	return id
}
