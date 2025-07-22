package sender

import (
	"log"
	"net/http"
	"strings"
)

func Msg(url, content string) error {
	data := strings.NewReader(content)
	res, err := http.Post(url, "application/x-www-form-urlencoded", data)
	userAgent := res.Header.Get("User-Agent")
	res.Header.Set("User-Agent", userAgent)
	if err != nil {
		log.Fatal(err)
	}
	
	return nil
}
