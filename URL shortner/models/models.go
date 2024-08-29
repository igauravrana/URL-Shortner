package models

import (
	"database/sql"
	"log"
	"time"

	"github.com/igauravrana/URL-Shortner/dbconnection"
)

type UrlData struct {
	Id          int       `json:"id"`
	OriginalUrl string    `json:"original_url"`
	ShortedUrl  string    `json:"shorted_url"`
	CreatedAt   time.Time `json:"created_at"`
}

var DB *sql.DB

func init() {
	DB = dbconnection.DB
}

// CreateURL inserts a new URL record into the database and returns the ID
func CreateURL(u UrlData) (int, error) {
	query := `INSERT INTO urldata (original_url, shorted_url, created_at) 
    VALUES ($1, $2, $3) RETURNING id`

	var id int
	err := dbconnection.DB.QueryRow(query, u.OriginalUrl, u.ShortedUrl, u.CreatedAt).Scan(&id)
	if err != nil {
		log.Println("Error creating URL:", err)
		return 0, err
	}
	return id, nil
}

// GetURLByID retrieves a URL record by ID
func GetURLByID(id int) (UrlData, error) {
	var urlData UrlData
	query := `SELECT id, original_url, shorted_url, created_at FROM urldata WHERE id = $1`
	row := dbconnection.DB.QueryRow(query, id)
	err := row.Scan(&urlData.Id, &urlData.OriginalUrl, &urlData.ShortedUrl, &urlData.CreatedAt)
	if err != nil {
		return urlData, err
	}
	return urlData, nil
}

// DeleteURL deletes a URL record by ID
func DeleteURL(id int) error {
	query := `DELETE FROM urldata WHERE id = $1`
	_, err := dbconnection.DB.Exec(query, id)
	return err
}

// GetURLByShortURL retrieves a URL record by its shortened URL
func GetURLByShortURL(shortUrl string) (UrlData, error) {
	var urlData UrlData
	query := `SELECT id, original_url, shorted_url, created_at FROM urldata WHERE shorted_url = $1`
	row := dbconnection.DB.QueryRow(query, shortUrl)
	err := row.Scan(&urlData.Id, &urlData.OriginalUrl, &urlData.ShortedUrl, &urlData.CreatedAt)
	if err != nil {
		return urlData, err
	}
	return urlData, nil
}
