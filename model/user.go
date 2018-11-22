package model

import "time"

// User is a user model
type User struct {
	ID                   uint      `json:"id" gorm:"column:ID"`
	UserName             string    `json:"UserName"  gorm:"column:UserName"`
	UserEmail            string    `json:"UserEmail"  gorm:"column:UserEmail"`
	UserPassword         string    `json:"UserPassword"   gorm:"column:UserPassword"`
	UserFirstName        string    `json:"UserFirstName" gorm:"column:UserFirstName"`
	UserLastName         string    `json:"UserLastName" gorm:"column:UserLastName"`
	UserCity             string    `json:"UserCity"   gorm:"column:UserCity"`
	UserState            string    `json:"UserState"   gorm:"column:UserState"`
	Deleted              bool      `json:"Deleted"   gorm:"column:Deleted"`
	UserRegistrationDate time.Time `json:"UserRegistrationDate"   gorm:"column:UserRegistrationDate"`
}

// PublicUser struct.
type PublicUser struct {
	*User
	Password omit `json:"password,omitempty"`
}

// SimpleUser is struct for retrieve user
type SimpleUser struct {
	UserName      string `json:"UserName"`
	UserFirstName string `json:"UserFirstName" description:"User firstname"`
	UserLastName  string `json:"UserLastName" description:"User lastname"`
}

// TableName indicates table name of user
func (User) TableName() string {
	return "user"
}
