package models

import (
	"realtime/app/auth"
	"time"
)

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Created     time.Time `json:"created" gorm:"autoCreateTime"`
	Updated     time.Time `json:"updated" gorm:"autoUpdateTime"`
	Created_UID uint      `json:"created_uid"`
	Updated_UID uint      `json:"updated_uid"`
	UserName    string    `json:"username"`
	Password    string    `json:"password"`
	Phone       string    `json:"phone"`
	MainPhoto   string    `json:"main_photo"`
}

func (mb *User) ToResponse() auth.UserResponseDTO {
	return auth.UserResponseDTO{
		ID:        mb.ID,
		Name:      mb.UserName,
		Password:  mb.Password,
		Phone:     mb.Phone,
		MainPhoto: mb.MainPhoto,
	}
}

type Message struct {
	ID         uint   `json:"id"`
	SenderID   uint   `json:"sender"`
	ReceiverID uint   `json:"receiver"`
	Text       string `json:"text"`
}
