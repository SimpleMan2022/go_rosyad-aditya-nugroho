package model_post_category_jwt

type Category struct {
	Id   int    `gorm:"primaryKey;type:int(11);autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(100)" json:"name"`
}

type CategoryRequest struct {
	Name string `json:"name"`
}

type CategoryResponse struct {
	Status     bool   `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
