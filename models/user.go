package models

import "github.com/jinzhu/gorm"

type user struct {
	ID       string `gorm:"primary_key" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
}

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string) (bool, error, string) {
	var usr user
	err := db.Select("id").Where(user{Username: username, Password: password}).First(&usr).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err, ""
	}

	if usr.ID != "" {
		return true, nil, usr.ID
	}

	return false, nil, ""
}
