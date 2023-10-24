package mysql

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	ID   uint64 `gorm:"id"`
	Rank int
	// MemberNumber string `gorm:"test"`
	Name string
	Age  uint8
	// Email     string
	// Birthday  *time.Time
	// ActivedAt sql.NullTime
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (u *User) TableName() string {
	return "users"
}

// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
//     fmt.Printf("before insert user id:%d\n", u.ID)
//     return
// }
//
// func (u *User) AfterCreate(tx *gorm.DB) (err error) {
//     fmt.Printf("after inser user id:%d\n", u.ID)
//     return
// }
