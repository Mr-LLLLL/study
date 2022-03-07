package main

import (
	"fmt"
	"gorm_test/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func main() {
	var err error
	// db, err = gorm.Open("mysql", "root:root@/test?charset=utf8&parseTime=True&loc=Local")
	db, err = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/test"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// createTable()
	// insert()
	find()
}

func find() {
	users := make([]model.User, 0)
	db.Model(&model.User{}).Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}

	user := new(model.User)
	user.ID = 128
	err := db.First(user).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("recored not found")
		return
	}

	fmt.Println(user)
}

func createTable() {
	// if !db.HasTable(&model.User{}) {
	//     db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.User{})
	// }
}

func insert() {
	m := model.User{
		// Model: g1.Model{
		//     // ID:        0,
		//     CreatedAt: time.Time{},
		//     UpdatedAt: time.Time{},
		// },
		ID:           122,
		MemberNumber: "",
		Name:         "",
		Age:          0,
		Email:        "",
		// Birthday:     &time.Time{},
	}

	users := []model.User{m, {Name: "Mr.w"}}
	// for _, user := range users {
	db.Debug().Create(users)
	// }
	// db.Create(&users)
}
