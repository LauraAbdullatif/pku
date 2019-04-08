package main

import (
	"encoding/json"
	"log"
	"net/http"
)

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//User is a struct
type User struct {
	CategoryID  int     `json:"category_id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Energy      int     `json:"energy,omitempty"`
	Protein     float32 `json:"protein,omitempty"`
	Phenylanine float32 `json:"phenylanine,omitempty"`
}

func users(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GO MYSQL")

	db, err := sql.Open("mysql", "root:pass1@tcp(127.0.0.1:3306)/pku")
	if err != nil {
		fmt.Println("err in db", err)
	}

	rows, err := db.Query("SELECT * FROM search")
	if err != nil {
		fmt.Println("NO rows", err)
		return
	}

	var users []User

	for rows.Next() {

		var CategoryID int
		var Name string
		var Energy int
		var Protein float32
		var Phenylanine float32

		err := rows.Scan(&CategoryID, &Name, &Energy, &Protein, &Phenylanine)
		users := append(users, User{CategoryID, Name, Energy, Protein, Phenylanine})

		if err != nil {
			fmt.Println("rows failed", err)
			return
		}

		usersBytes, _ := json.Marshal(&users)

		w.Write(usersBytes)
		defer db.Close()
	}
}
func main() {
	http.HandleFunc("/", users)
	log.Fatal(http.ListenAndServe(":8083", nil))

}
