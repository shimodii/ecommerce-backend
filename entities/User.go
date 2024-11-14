package entities

type User struct {
    ID          uint    `json:"id" gorm:"primaryKey"`
    Name        string  `json:"name"`
    Email       string  `json:"email" gorm:"uique"`
    Password    string  `json:"_"`
    Role        string  `json:"role"`
}
