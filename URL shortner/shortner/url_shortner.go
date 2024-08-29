package shortner

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/igauravrana/URL-Shortner/models"
)

// GenerateShortURL creates a shortened version of the original URL
func GenerateShortURL(urlToBeShort models.UrlData) string {
	hasher := md5.New()
	hasher.Write([]byte(urlToBeShort.OriginalUrl))
	data := hasher.Sum(nil)

	encodedIntoStr := hex.EncodeToString(data)
	return encodedIntoStr[:8]
}
