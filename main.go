package main

import (
	"database/sql"
	"fmt"

	db "./db"
)

var data *sql.DB

//open database
func init() {
	var err error
	data, err = sql.Open("postgres", "user=DSC password=DSC sslmode=disable dbname=database")
	if err != nil {
		panic(err)
	}
	err = data.Ping()
	if err != nil {
		panic(err)
	}
}
func main() {

	var val db.Phone
	val = db.Phone{
		Name:   "Uti Michael",
		Number: "08115780877",
	}
	val.Create(data)
	fmt.Println(val)
	val.Delete(data)
}
