package db

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"go-is-cool/src/utils"
)

type UserEntity struct {
	id   uuid.UUID
	name string
}

func PersistUser(name string) string {
	config := utils.GetConfig()
	user := UserEntity{uuid.New(), name}

	connStr := fmt.Sprintf("postgresql://%s:%s@localhost:%s/postgres?sslmode=disable", config.Db.User, config.Db.Password, config.Db.Port)
	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	if err != nil {
		fmt.Println("Error! ", err)
	} else {
		_, err := db.Exec("INSERT INTO users VALUES ($1, $2)", user.id.String(), user.name)
		if err != nil {
			fmt.Println("Error! ", err)
		}
	}
	return user.id.String()
}
