package model_post

type Comment struct {
	Id      int    `gorm:"primaryKey;autoIncrement;type:int(11)" json:"id"`
	PostId  int    `gorm:"not null" json:"post_id"`
	Content string `gorm:"type:text" json:"content"`
	Post    *Post  `gorm:"foreignKey:PostId" json:"post,omitempty"`
}

type CommentRequest struct {
	Content string `json:"content"`
}

type CommentResponse struct {
	Status     bool   `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
