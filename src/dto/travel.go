package dto

type TravelDTO struct {
	Destination string  `json:"destination"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	Price       float64 `json:"price"`
	Type        string  `json:"type"`
	Photo       []byte  `json:"image" gorm:"type:blob"`
	UserID      uint    `json:"userId"`
}
