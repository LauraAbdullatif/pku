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

//Nutrients is a struct
type Nutrients struct {
	NutrientsID int     `json:"nutrients_id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Energy      int     `json:"energy,omitempty"`
	Protein     float32 `json:"protein,omitempty"`
	Phenylanine float32 `json:"phenylanine,omitempty"`
}

func allFoods(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GO MYSQL")

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/pku")
	if err != nil {
		fmt.Println("err in db", err)
	}

	rows, err := db.Query("SELECT * FROM FOOD_TABLE")
	if err != nil {
		fmt.Println("NO rows", err)
		return
	}

	var allFoods []Nutrients

	for rows.Next() {

		var NutrientsID int
		var Name string
		var Energy int
		var Protein float32
		var Phenylanine float32

		err := rows.Scan(&NutrientsID, &Name, &Energy, &Protein, &Phenylanine)
		allFoods := append(allFoods, Nutrients{NutrientsID, Name, Energy, Protein, Phenylanine})

		if err != nil {
			fmt.Println("rows failed", err)
			return
		}

		allFoodsBytes, _ := json.Marshal(&allFoods)

		w.Write(allFoodsBytes)
		defer db.Close()
	}
}
func main() {
	http.HandleFunc("/", allFoods)
	log.Fatal(http.ListenAndServe(":8083", nil))

}
