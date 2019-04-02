package database_handler

import (
	"fmt"
)

// struct to hold data from postgres
type Entry struct {
	Bin            string `json:"bin"`
	Cnstrct_yr     string `json:"cnstrct_yr"`
	Lstmoddate     string `json:"lstmoddate"`
	Lststatype     string `json:"lststatype"`
	Doitt_id       string `json:"doitt_id"`
	Heightroof     string `json:"heightroof"`
	Feat_code      string `json:"feat_code"`
	Groundelev     string `json:"groundelev"`
  Shape_area     string `json:"shape_area"`
  Shape_len      string `json:"shape_len"`
  Base_bbl       string `json:"base_bbl"`
  Mpluto_bbl     string `json:"mpluto_bbl"`
  Geomsource     string `json:"geomsource"`
}

func GetEntry(id uint) (*Entry) {
	entry := &Entry{}
	err := GetDB().Table("ny_data").Where("id = ?", id).First(entry).Error
	if err != nil {
		return nil
	}
	return entry
}

func GetAll() ([]*Entry) {
	entries := make([]*Entry, 0)
	err := GetDB().Table("ny_data").Find(&entries).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return entries
}

func GetAfterYear(year []string) ([]*Entry) {
	entries := make([]*Entry, 0)
	err := GetDB().Table("ny_data").Where("cnstrct_yr > ?", year).Find(&entries).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return entries
}
