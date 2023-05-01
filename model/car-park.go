package model

type CarPark struct {
	ID           uint    `gorm:"primaryKey"`
	Name         string  `json:"name"`
	City         string  `json:"city"`
	Country      string  `json:"country"`
	Zipcode      string  `json:"zipcode"`
	Latitude     float32 `json:"latitude"`
	Longitude    float32 `json:"longitude"`
	Description  string  `json:"description"`
	TotalSpaces  int32   `json:"total_spaces"`
	PricePerHour float32 `json:"price_per_hour"`
	ParkingSpots []ParkingSpot
    Reservations []Reservation
}
