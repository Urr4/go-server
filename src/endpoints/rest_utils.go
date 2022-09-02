package endpoints

import (
	"errors"
	"net/http"
	"strings"
)

func GetAuthToken(r *http.Request) (e error, token string) {
	for key, values := range r.Header {
		if key == "Authorization" {
			for _, value := range values {
				if strings.Contains(value, "Bearer") {
					token = strings.TrimPrefix(value, "Bearer ")
					return
				}
			}
		}
	}
	e = errors.New("missing authorization header")
	return
}