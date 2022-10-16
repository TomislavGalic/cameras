package models

type Camera struct {
	Id           uint    `json:"-" gorm:"primaryKey"`
	LocationName string  `json:"locationName"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}
