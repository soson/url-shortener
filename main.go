package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"

	"url-shortener/hash"
)

var db *sql.DB

// URL struct for keep url data in database
type URL struct {
	ID       int    `json:"id,omitempty"`
	ShortURL string `json:"short_url,omitempty"`
	LongURL  string `json:"long_url,omitempty"`
}

// CreateURL add short url to database
func CreateURL(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	stmt, err := db.Prepare("INSERT INTO url (short_url, long_url) VALUES (?, ?)")
	defer stmt.Close()

	if err != nil {
		panic(err.Error()) // TODO: proper error handling instead of panic in your app
	}

	var url URL
	_ = json.NewDecoder(req.Body).Decode(&url)
	url.ShortURL = hash.RandStringBytes(8)

	_, err = stmt.Exec(url.ShortURL, url.LongURL)
	if err != nil {
		panic(err.Error()) // TODO: proper error handling instead of panic in your app
	}

	json.NewEncoder(w).Encode(url)
}

// UpdateURL edit url in database
func UpdateURL(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	stmt, err := db.Prepare("UPDATE url SET long_url=? WHERE short_url=?")
	defer stmt.Close()

	if err != nil {
		panic(err.Error()) // TODO: proper error handling instead of panic in your app
	}

	var url URL
	_ = json.NewDecoder(req.Body).Decode(&url)
	url.ShortURL = ps.ByName("url")

	_, err = stmt.Exec(url.LongURL, url.ShortURL)
	if err != nil {
		panic(err.Error()) // TODO: proper error handling instead of panic in your app
	}

	json.NewEncoder(w).Encode(url)
}

// DeleteURL edit url in database
func DeleteURL(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	stmt, err := db.Prepare("DELETE FROM url WHERE short_url=?")
	defer stmt.Close()

	if err != nil {
		panic(err.Error()) // TODO: proper error handling instead of panic in your app
	}

	_, err = stmt.Exec(ps.ByName("url"))
	if err != nil {
		panic(err.Error()) // TODO: proper error handling instead of panic in your app
	}

	w.WriteHeader(204)
}

// ReadURL serves short urls
func ReadURL(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	stmt, err := db.Prepare("SELECT short_url, long_url FROM url WHERE short_url=?")
	defer stmt.Close()

	if err != nil {
		panic(err.Error()) // TODO: proper error handling instead of panic in your app
	}

	var url URL

	err = stmt.QueryRow(ps.ByName("url")).Scan(&url.ShortURL, &url.LongURL)
	if err != nil {
		panic(err.Error()) // TODO: proper error handling instead of panic in your app
	}

	json.NewEncoder(w).Encode(url)
}

func main() {
	db, _ = sql.Open("mysql", "user:password@/database")
	defer db.Close()

	// Validate DSN data:
	err := db.Ping()
	if err != nil {
		panic(err.Error()) // TODO: proper error handling instead of panic in your app
	}

	router := httprouter.New()
	router.POST("/", CreateURL)
	router.PUT("/:url", UpdateURL)
	router.DELETE("/:url", DeleteURL)
	router.GET("/:url", ReadURL)

	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
