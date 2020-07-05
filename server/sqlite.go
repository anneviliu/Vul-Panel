package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

func (s *Service) initDb() {
	db, err := gorm.Open("sqlite3", fmt.Sprintf("./database.db",
	))

	if err != nil {
		log.Fatalln(err)
	}

	s.Db = db
	s.Db.AutoMigrate(&Vul{}, &Pushed{}, &RegUser{})
}
