package endpoints

import (
	"errors"
	"fmt"
	"go-is-cool/src/db"
	rest_client "go-is-cool/src/rest-client"
	"go-is-cool/src/utils"
	"net/http"
)

func Greet(responseWriter http.ResponseWriter, request *http.Request) {

	config := utils.GetConfig()
	_, token := GetAuthToken(request)
	if config.Auth.Password != token {
		responseWriter.WriteHeader(403)
		responseWriter.Write([]byte("Nice try hackerboy"))
		return
	}
	fmt.Println("Got token",token)
	_, name := getName(request)
	id := db.PersistUser(name)
	username := rest_client.GetRandomUsername()
	responseWriter.Write([]byte("Hello "+name+". Your ID is "+id+". Do you know "+username+"?"))
}

func getName(request *http.Request) (e error, name string) {
	for key, queryParam := range request.URL.Query() {
		if key == "name" {
			name = queryParam[0]
			return
		}
	}
	e = errors.New("missing query-param 'name'")
	return
}