package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//Phone ....
//due to golang taking numbers starting with 0 and since some numbers might be added as +234 i'll make number a string
type Phone struct {
	Name   string
	Number string
	ID     int
}

/*Create data returning error if failed*/
func (data *Phone) Create(db *sql.DB) error {
	statement := `
	insert into phone (author, accnumber)
	values ($1, $2)
	returning id`
	stmt, err := db.Prepare(statement)
	if err != nil {
		fmt.Println("err creating", err)
		return err
	}
	err = stmt.QueryRow(data.Name, data.Number).Scan(&data.ID)
	return err
}

//Delete ...
/*deletes content in database*/
func (data *Phone) Delete(db *sql.DB) error {

	_, err := db.Exec("delete from phone where id = $1", data.ID)
	return err
}

func (data *Phone) Update(db *sql.DB) error {
	_, err := db.Exec("update phone set number = $2,name = $1 where id = $1", data.ID, data.Name, data.Number)
	return err
}
