package rest_client

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"strconv"
)

type UserResponse struct {
	Id   int    `json:id`
	Name string `json:name`
}

func GetRandomUsername() (username string) {

	id := rand.Intn(10)

	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/users", nil)
	q := req.URL.Query()
	q.Add("id", strconv.Itoa(id))
	req.URL.RawQuery = q.Encode()

	response, err := http.Get(req.URL.String())
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "Reading bytes failed"
	}
	users := make([]UserResponse, 1)
	json.Unmarshal(respBytes, &users)
	if len(users) != 0 {
		username = users[0].Name
	} else {
		username = "Nobody"
	}
	return
}
