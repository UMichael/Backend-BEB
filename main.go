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
	fmt.Println(val)
	val.Create(data)
	value, err := Retrieve(val.Number)
	if err != nil {
		fmt.Println("There's no such data", err)
	}
	val.Number = "08115780876"
	err = val.Update(data)
	if err != nil {
		fmt.Println("Cant update", err)
	}
	value, err = Retrieve(val.Number)
	if err != nil {
		fmt.Println("There's no such data", err)
	}
	fmt.Println(value)
	val.Delete(data)
}

/*used user number to get details*/
func Retrieve(number string) (value db.Phone, err error) {
	value = db.Phone{}
	err = data.QueryRow("select accnumber, author from phone where accnumber = $1", number).Scan(&value.Number, &value.Name)
	return
}
