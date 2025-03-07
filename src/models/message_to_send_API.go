package models

type SendMessage struct {
    UserId     int32  `json:"user_id"`
    FullName   string `json:"full_name"`
    Email      string `json:"email"`
    Description string `json:"description"`
}