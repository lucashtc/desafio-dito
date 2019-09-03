package coletora

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
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
func Insert(con *gorm.DB,values Colector) (error){
	con.Create(&values)
	if con.Error != nil {
		return errors.Wrap(con.Error,"falha ao fazer insert")
	}
	defer con.Close()
	return nil;
}