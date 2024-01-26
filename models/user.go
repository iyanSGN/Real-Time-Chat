package models

import "time"

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
	Created_UID uint      `json:"created_uid"`
	Updated_UID uint      `json:"updated_uid"`
	UserName    string    `json:"username"`
	Password    string    `json:"password"`
	Phone       string    `json:"phone"`
}

type Message struct {
	ID         uint   `json:"id"`
	SenderID   uint   `json:"sender"`
	ReceiverID uint   `json:"receiver"`
	Text       string `json:"text"`
}


