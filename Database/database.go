package database

import (
	"encoding/json"
	"log"
	"os"
)

var Users []User

func InitDataBase() {
	LoadUsers()
	
}

func LoadUsers() {
	userfile := "users.json"

	nByte, err := os.ReadFile(userfile)
	if err != nil {
		log.Println("Can not read the file due to :", err, " Loading with No Active user!")
	}

	err = json.Unmarshal(nByte, &Users)
	if err != nil {
		log.Println(err)
	}
}
