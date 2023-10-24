package pg

import (
	"time"
)

type ResultComponent struct {
	ID              string    `gorm:"id"`
	TaskId          int       `gorm:"task_id"`
	Name            string    `gorm:"name"`
	RootComponentId string    `gorm:"root_component_id"`
	SecurityLevelId int       `gorm:"security_level_id"`
	CreateDate      time.Time `gorm:"create_date"`
	ModifyDate      time.Time `gorm:"modify_date"`
}

func (u *ResultComponent) TableName() string {
	return "result_component"
}
