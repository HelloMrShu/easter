package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	var err error
	db := ServerConfig.Db

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", db.Username, db.Password, db.Host, db.Port, db.Database)
	fmt.Println(dsn)

	EasterDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if nil != err {
		panic(err)
	}
}
