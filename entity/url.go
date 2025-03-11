package entity

import "time"

type ShortenURLRequest struct {
	OriginalURL string `json:"original_url" validate:"required" example:"https://www.youtube.com/watch?v=sXdIVfeId9E&ab_channel=AsditaPrasetya"`
	Expiry      string `json:"expiry,omitempty" validate:"omitempty,date" example:"2025-05-12"`
}

type ShortenURLResponse struct {
	ShortCode string     `json:"short_code" example:"C2K7fG"`
	ShortURL  string     `json:"short_url" example:"http://localhost:3000/api/v1/61PjEp"`
	Expiry    *time.Time `json:"expiry,omitempty" example:"2025-05-12T23:59:59Z"`
}

type RedirectResponse struct {
	OriginalURL string `json:"original_url" example:"https://www.youtube.com/watch?v=sXdIVfeId9E&ab_channel=AsditaPrasetya"`
}
