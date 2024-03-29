package model_fruits

import "github.com/google/uuid"

type FruitApi struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Nutritions struct {
		Id            int     `json:"id"`
		Calories      float64 `json:"calories"`
		Fat           float64 `json:"fat"`
		Sugar         float64 `json:"sugar"`
		Carbohydrates float64 `json:"carbohydrates"`
		Protein       float64 `json:"protein"`
	} `json:"nutritions"`
}

type FruitRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Nutritions struct {
	Id            int     `json:"id"`
	Calories      float64 `json:"calories"`
	Fat           float64 `json:"fat"`
	Sugar         float64 `json:"sugar"`
	Carbohydrates float64 `json:"carbohydrates"`
	Protein       float64 `json:"protein"`
}

type Fruit struct {
	Id          uuid.UUID  `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"type:varchar(100)" json:"name"`
	Price       int        `json:"price"`
	NutritionId int        `gorm:"not null" json:"nutrition_id"`
	Nutrition   Nutritions `gorm:"foreignKey:NutritionId"`
}

type FruitResponse struct {
	Status     bool   `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
