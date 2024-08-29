package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/igauravrana/URL-Shortner/models"
	"github.com/igauravrana/URL-Shortner/shortner"
)

func CreateShortURL(originalUrl string) string {
	newShortedUrlData := models.UrlData{
		OriginalUrl: originalUrl,
		ShortedUrl:  shortner.GenerateShortURL(models.UrlData{OriginalUrl: originalUrl}),
		CreatedAt:   time.Now(),
	}

	// Save to database
	id, err := models.CreateURL(newShortedUrlData)
	if err != nil {
		fmt.Println("Error saving URL:", err)
		return ""
	}
	fmt.Println("Short URL created with ID:", id)

	// Return the shortened URL
	return newShortedUrlData.ShortedUrl
}

// ShortenURLHandler handles the shortening of URLs
func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Url string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	shortUrl := CreateShortURL(data.Url)
	if shortUrl == "" {
		http.Error(w, "Failed to create short URL", http.StatusInternalServerError)
		return
	}

	response := struct {
		ResponseUrl string `json:"shorted_url"`
	}{ResponseUrl: shortUrl}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Unable to give shorted url :", http.StatusBadRequest)
	}
}

// RedirectToOriginalUrl redirects to the original URL based on the shortened URL
func RedirectToOriginalUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortUrl := vars["shortUrl"]

	url, err := models.GetURLByShortURL(shortUrl)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, url.OriginalUrl, http.StatusFound)
}

// DeleteURLHandler handles the deletion of a URL by its shortened version
func DeleteURLHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortUrl := vars["shortUrl"]

	url, err := models.GetURLByShortURL(shortUrl)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	if err := models.DeleteURL(url.Id); err != nil {
		http.Error(w, "Failed to delete URL", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetURLHandler retrieves a URL by its shortened version
func GetURLHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortUrl := vars["shortUrl"]

	urlData, err := models.GetURLByShortURL(shortUrl)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(urlData); err != nil {
		http.Error(w, "Unable to encode URL data", http.StatusInternalServerError)
	}
}
