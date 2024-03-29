package model_post

type Post struct {
	Id      int       `gorm:"primaryKey;type:int(11);autoIncrement" json:"id"`
	Title   string    `gorm:"type:varchar(100)" json:"title"`
	Content string    `gorm:"type:text" json:"content"`
	Comment []Comment `gorm:"foreignKey:PostId;constraint:OnDelete:CASCADE" json:"comment"`
}

type PostRequest struct {
	Title   string `gorm:"type:varchar(100)" json:"title"`
	Content string `gorm:"type:text" json:"content"`
}

type PostResponse struct {
	Status     bool   `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
