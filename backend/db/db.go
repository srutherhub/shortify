package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	dbsqlc "shortify/db/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/teris-io/shortid"
)

var Queries *dbsqlc.Queries
var ctx = context.Background()

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading environment variables")
	}

	dbURL := os.Getenv("DB_URL")

	fmt.Println(os.Getenv("DB_URL"))
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		panic(err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}
	Queries = dbsqlc.New(pool)
}

func GetURL(id string) (dbsqlc.GetUrlFromIDRow, error) {
	var data dbsqlc.GetUrlFromIDRow
	data, err := Queries.GetUrlFromID(ctx, id)
	if err != nil {
		return dbsqlc.GetUrlFromIDRow{Url: "", RouteName: ""}, err
	}
	return data, nil
}

func InsertNewURL(url string) (string, error) {
	id := CreateUniqueID()
	now := time.Now().Unix()

	if len(url) < 8 {
		return "", errors.New("invalid url")
	}

	if !isURLHttps(url) {
		return "", errors.New("url is not https")
	}

	err := Queries.InsertNewUrl(ctx, dbsqlc.InsertNewUrlParams{ID: id, Url: url, CreatedAt: now})

	if err != nil {
		return "", err
	}
	return id, nil
}

func IncrementURLClickCount(id string) error {
	err := Queries.IncrementUrlClickCount(ctx, id)
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

func IsUrlActive(created_at int64, expires_at pgtype.Int8) bool {
	now := time.Now().Unix()
	const default180DayExpiry = 180 * 24 * 60 * 60
	if expires_at.Valid {
		if expires_at.Int64 <= now {
			return false
		} else {
			return true
		}
	}

	if now >= created_at+default180DayExpiry {
		return false
	}
	return true
}
