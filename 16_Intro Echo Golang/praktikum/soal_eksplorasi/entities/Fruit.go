package entities

import "github.com/google/uuid"

type Fruit struct {
	Name       string `json:"name"`
	Nutritions struct {
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
	Calories      float64 `json:"calories"`
	Fat           float64 `json:"fat"`
	Sugar         float64 `json:"sugar"`
	Carbohydrates float64 `json:"carbohydrates"`
	Protein       float64 `json:"protein"`
}

type Data struct {
	Id         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	Price      int        `json:"price"`
	Nutritions Nutritions `json:"nutritions"`
}

type FruitResponse struct {
	Status     bool   `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
