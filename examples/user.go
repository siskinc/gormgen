package example

import "github.com/jinzhu/gorm"

var client *gorm.DB

//go:generate gormgen -structs User -output user_gen.go -client client
type User struct {
	gorm.Model
	Name  string
	Age   int
	Email string
}
