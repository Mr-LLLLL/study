package pg

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var pgdb *gorm.DB

func Pg_op() {
	var err error
	// pgdb, err = gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"), &gorm.Config{})
	pgdb, err = gorm.Open(postgres.Open("host=10.1.2.38 user=gpadmin password=Anpr0123465^_^ dbname=sca_web port=5432 sslmode=disable TimeZone=Asia/Shanghai"), &gorm.Config{})
	if err != nil || pgdb == nil {
		log.Fatal(err)
	}

	find()
}

func find() {
	users := make([]ResultComponent, 0)
	err := pgdb.Model(&ResultComponent{}).Limit(10).Find(&users).Debug().Error
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range users {
		fmt.Println(v.CreateDate.Format("2006-01-02 15:04:05"))
	}
}
