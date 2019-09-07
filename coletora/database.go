package coletora

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	
	// drive mysql for connection 
	_"github.com/jinzhu/gorm/dialects/mysql"
)

// Conn return connection instance.
func Conn() (*gorm.DB,error){
	db, err := gorm.Open("mysql", "root@/dito?charset=utf8")
	if err != nil {
		return nil,err
	}
	return db,err
}

// Insert ...
func Insert(con *gorm.DB,values Event) (error){
	con.Create(&values)
	if con.Error != nil {
		return errors.Wrap(con.Error,"falha ao fazer insert")
	}
	defer con.Close()
	return nil;
}

// Get ...
func Get(con *gorm.DB,name string) ([]Event, error){
	events := []Event{}
	name = fmt.Sprint("%" + name + "%")
	con.Where("NAME LIKE ?",name ).Select("DISTINCT(NAME)").Find(&events)
	if con.Error != nil {
		return nil,con.Error
	}
	return events,nil
}