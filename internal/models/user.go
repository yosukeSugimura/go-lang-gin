package models

import (
	"gorm.io/gorm"
)

// User構造体はユーザー情報を保持します
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"size:100;not null"`
	Email    string `json:"email" gorm:"uniqueIndex;size:100;not null"`
	Password string `json:"-" gorm:"size:255;not null"` // JSONレスポンスには表示しない
}

// CreateUserは新しいユーザーをデータベースに追加します
func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

// GetUserByIDは指定されたIDでユーザーを検索します
func GetUserByID(db *gorm.DB, id uint) (*User, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUserは既存のユーザー情報を更新します
func UpdateUser(db *gorm.DB, user *User) error {
	return db.Save(user).Error
}

// DeleteUserは指定されたIDのユーザーを削除します
func DeleteUser(db *gorm.DB, id uint) error {
	return db.Delete(&User{}, id).Error
}
