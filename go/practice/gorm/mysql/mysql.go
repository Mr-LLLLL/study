package mysql

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
)

var mysqldb *gorm.DB

func mysql_op() {
	var err error
	// db, err = gorm.Open("mysql", "root:root@/test?charset=utf8&parseTime=True&loc=Local")
	mysqldb, err = gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/test"), &gorm.Config{})
	if err != nil {

		log.Fatal(err)
	}

	// createTable()
	insert()
	// find()
}

func find() {
	users := make([]User, 0)
	mysqldb.Model(&User{}).Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}

	user := new(User)
	user.ID = 128
	err := mysqldb.First(user).Error
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
	// db.Create([]*model.User{{ID: 1, Rank: 1}, {ID: 2, Rank: 2}})

	wg := sync.WaitGroup{}
	ch := make(chan struct{}, 100)
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		i := i
		ch <- struct{}{}
		go func() {
			defer func() {
				wg.Done()
				<-ch
			}()

			start := 1000 * i
			end := 1000 * (i + 1)
			users := make([]*User, 0, 1000)
			for i := start; i < end; i++ {
				users = append(users, &User{
					ID:   uint64(i),
					Rank: i,
					// Age: int(i),
				})
			}
			err := mysqldb.Create(users).Error
			if err != nil {
				fmt.Println(err)
				return
			}
		}()
	}
	wg.Wait()
}
