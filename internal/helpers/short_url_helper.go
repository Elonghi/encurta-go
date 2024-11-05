package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"net/url"
	"regexp"
	"strings"
)

func GenerateShortURL(url string) string {
	randomBytes := make([]byte, 4)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	encodedUrl := base64.URLEncoding.EncodeToString(randomBytes)[:6]
	return sanitizeUrl(encodedUrl)
}

func sanitizeUrl(url string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9-_.~%]+`)
	return strings.ToLower(re.ReplaceAllString(url, ""))
}

func CheckURL(testUrl string) bool {
	parsedUrl, err := url.ParseRequestURI(testUrl)
	if err != nil {
		return false
	}
	if parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
		return false
	}

	return true
}
