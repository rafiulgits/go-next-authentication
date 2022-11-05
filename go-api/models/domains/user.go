package domains

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID              int    `gorm:"primarykey"`
	Name            string `gorm:"type:varchar(256);not null"`
	Email           string `gorm:"type:varchar(256);not null;unique"`
	Phone           string `gorm:"type:varchar(20);not null;unique"`
	AuthProvider    string `gorm:"type:varchar(20);not null"`
	Password        string `gorm:"type:varchar(512);not null"`
	IsEmailVerified bool   `gorm:"type:bool"`
}

const UserTableName = "Users"

func (User) TableName() string {
	return UserTableName
}

func (user *User) SetPassword(rawPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) CheckIfPasswordIsCorrect(rawPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(rawPassword)); err != nil {
		return false
	}
	return true
}
