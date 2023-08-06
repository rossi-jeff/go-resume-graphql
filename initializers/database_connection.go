package initializers

import (
	"fmt"
	"log"

	"github.com/rossi-jeff/go-resume-graphql/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnect() {
	conf := config.New()
	dsn := ConnectionString(conf.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("database connected")
	}
	DB = db
}

func ConnectionString(d config.DbConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", d.DbUser, d.DbPass, d.DbHost, d.DbPort, d.DbName)
}
