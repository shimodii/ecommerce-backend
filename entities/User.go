package entities

type User struct {
    ID      uint    `json:"id" gorm:"primaryKey"`
}
