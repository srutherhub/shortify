package db

import (
	"context"
	"errors"
	"log"
	"os"
	dbsqlc "shortify/db/sqlc"
	"shortify/models"
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

func InsertNewURL(data models.ShortifyRequestBody, username string) (string, error) {
	id := CreateUniqueID()
	now := time.Now().Unix()

	if len(data.URL) < 8 {
		return "", errors.New("invalid url")
	}

	if !isURLHttps(data.URL) {
		return "", errors.New("url is not https")
	}

	ExpiresAt := pgtype.Int8{Int64: now + data.ExpiresIn, Valid: data.ExpiresIn != 0}
	UtmSource := pgtype.Text{String: data.UtmSource, Valid: data.UtmSource != ""}
	UtmMedium := pgtype.Text{String: data.UtmMedium, Valid: data.UtmMedium != ""}
	UtmCampaign := pgtype.Text{String: data.UtmMedium, Valid: data.UtmCampaign != ""}
	UtmContent := pgtype.Text{String: data.UtmContent, Valid: data.UtmContent != ""}
	UtmTerm := pgtype.Text{String: data.UtmTerm, Valid: data.UtmTerm != ""}

	err := Queries.InsertNewUrl(ctx, dbsqlc.InsertNewUrlParams{ID: id, Url: data.URL, CreatedAt: now, ExpiresAt: ExpiresAt, UtmSource: UtmSource, UtmMedium: UtmMedium, UtmCampaign: UtmCampaign, UtmContent: UtmContent, UtmTerm: UtmTerm, Username: username})

	if err != nil {
		return "", err
	}
	return id, nil
}

func IncrementURLClickCount(id string) error {
	err := Queries.IncrementUrlClickCount(ctx, id)
	return err
}

func DeleteExpireUrl(id string, route string) error {
	err := Queries.DeleteExpiredUrl(ctx, dbsqlc.DeleteExpiredUrlParams{ID: id, RouteName: route})
	if err != nil {
		return err
	}
	return nil
}

func GetUserUrls(username string) ([]dbsqlc.GetUserUrlsRow, error) {
	urls, err := Queries.GetUserUrls(ctx, username)

	if err != nil {
		return nil, err
	}
	return urls, nil
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
