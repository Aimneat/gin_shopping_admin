package models

import "time"

type Rights struct {
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	AuthName  string     `json:"authName"` //权限说明
	Level     string     `json:"level"`    //权限层级
	Pid       uint       `json:"pid"`      //父权限id
	Path      string     `json:"path"`     //对应访问路径
	Children  []Rights   `json:"children"`
	Roles     []Roles
}
