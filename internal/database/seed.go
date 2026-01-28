package database

import (
	"github.com/rynr00/go-resto/internal/model"
	"github.com/rynr00/go-resto/internal/model/constant"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(&model.MenuItem{}, &model.Order{}, &model.ProductOrder{}, &model.User{})

	foodMenu := []model.MenuItem{
		{
			Name:      "Mie Ayam",
			OrderCode: "Mie_ayam",
			Price:     13000,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Ayam Geprek",
			OrderCode: "ayam_geprek",
			Price:     12500,
			Type:      constant.MenuTypeFood,
		},
	}

	drinkMenu := []model.MenuItem{
		{
			Name:      "Es Teh Manis",
			OrderCode: "es_teh_manis",
			Price:     5000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Air Mineral",
			OrderCode: "air_mineral",
			Price:     4000,
			Type:      constant.MenuTypeDrink,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)  //insert data
		db.Create(&drinkMenu) //insert data
	}
}
