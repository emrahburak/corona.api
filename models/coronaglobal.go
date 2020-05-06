package models

import (
	"github.com/jinzhu/gorm"
)

type CoronaGlobal struct {
	gorm.Model
	Confirmed int `json:"confirmed"`
	Deaths    int `json:"deaths"`
	Recovered int `json:"recovered"`
}

type CoronaGlobalCountries struct {
	gorm.Model
	Confirmed   int     `json:"confirmed"`
	CountryCode string  `json:"country_code"`
	Dead        int     `json:"dead"`
	Latitude    float64 `json:"latitude"`
	Location    string  `json:"location"`
	Longitude   float64 `json:"longitude"`
	Recovered   int     `json:"recovered"`
	Updated     string  `json:"updated"`
}

type CoronaGlobalCountriesResponse struct {
	gorm.Model
	CountryList []CoronaGlobalCountries `json:"countryList"`
}

// type CoronaGlobalModel struct {
// 	gorm.Model

// 	Confirmed int `gorm:"not null; size:30"`
// 	Deaths    int `gorm:"not null; size:30"`
// 	Recovered int `gorm:"not null; size:30"`
// }
