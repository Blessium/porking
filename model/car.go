package model

type Car struct {
    ID uint `gorm:"primaryKey"`
    Name string `json:"name"`
    Size string `json:"size"`
    Color string `json:"color"`
    LicensePlate string `json:"license_plate"`
}
