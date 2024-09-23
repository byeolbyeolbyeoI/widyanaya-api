package model

import "time"

type Publication struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title" validate:"required,max=255"`
	CoverURL    string    `json:"cover_url" validate:"required,url,max=255"`
	Description string    `json:"description"`
	OpeningDate time.Time `json:"opening_date"`
	ClosingDate time.Time `json:"closing_date" validate:"gtefield=OpeningDate"`
	Volume      int       `json:"volume" validate:"min=0"`
	Year        int       `json:"year" validate:"required,min=1900,max=9999"`
	UserID      int64     `json:"user_id"`
	CategoryID  int64     `json:"category_id"`
}

type PublicationCategory struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description"`
}
