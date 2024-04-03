package model_post_category_jwt

type Post struct {
	Id         int      `gorm:"primaryKey;type:int(11);autoIncrement" json:"id"`
	Title      string   `gorm:"type:varchar(100)" json:"title"`
	Content    string   `gorm:"type:text" json:"content"`
	CategoryId int      `gorm:"type:int(11);not null" json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryId" json:"Category"`
}

type PostRequest struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	CategoryId int    `json:"category_id"`
}

type PostResponse struct {
	Status     bool   `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
