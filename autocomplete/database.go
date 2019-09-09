package autocomplete

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	// drive mysql for connection
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Conn return connection instance.
func Conn() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:password@tcp(db:3306)/dito?charset=utf8")
	if err != nil {
		return nil, err
	}
	return db, err
}

// Insert ...
func Insert(con *gorm.DB, values Event) error {
	con.Create(&values)
	if con.Error != nil {
		return errors.Wrap(con.Error, "falha ao fazer insert")
	}
	defer con.Close()
	return nil
}

// Get ...
func Get(con *gorm.DB, name string) ([]string, error) {
	events := []string{}
	name = fmt.Sprint("%" + name + "%")

	rows, err := con.Raw("select DISTINCT(EVENT) from events where EVENT LIKE ?", name).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&name)
		events = append(events, name)
	}
	return events, nil
}
