package model

import "time"

type URL struct {
	ID          int        `json:"id" gorm:"type:int;primary_key"`
	ShortCode   string     `json:"short_code" gorm:"uniqueIndex; not null"`
	OriginalURL string     `json:"original_url" gorm:"not null"`
	Clicks      int        `json:"clicks" gorm:"default:0"`
	ExpiredAt   *time.Time `json:"expired_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (URL) TableName() string {
	return "urls"
}
