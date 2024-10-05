package model

import "time"

type Publication struct {
	Id                    int       `json:"id"`
	Title                 string    `json:"title" validate:"required,max=255"`
	CoverURL              string    `json:"cover_url" validate:"required,url,max=255"`
	Description           string    `json:"description" validate:"required"`
	Volume                int       `json:"volume" validate:"required,min=0"`
	Year                  int       `json:"year" validate:"required,min=1000,max=9999"`
	OpeningDate           time.Time `json:"opening_date" validate:"required"`
	ClosingDate           time.Time `json:"closing_date" validate:"required"`
	ReviewEstimation      int       `json:"review_estimation"`
	PublisherId           int       `json:"publisher_id" validate:"required"`
	PublicationCategoryID int       `json:"publication_category_id" validate:"required"`
}

type UpdatedPublication struct {
	Id                    int       `json:"id" validate:"required"`
	Title                 string    `json:"title" validate:"required,max=255"`
	CoverURL              string    `json:"cover_url" validate:"required,url,max=255"`
	Description           string    `json:"description" validate:"required"`
	Volume                int       `json:"volume" validate:"required,min=0"`
	Year                  int       `json:"year" validate:"required,min=1000,max=9999"`
	OpeningDate           time.Time `json:"opening_date" validate:"required"`
	ClosingDate           time.Time `json:"closing_date" validate:"required"`
	ReviewEstimation      int       `json:"review_estimation"`
	PublisherId           int       `json:"publisher_id" validate:"required"`
	PublicationCategoryID int       `json:"publication_category_id" validate:"required"`
}

type PublicationCategory struct {
	Id          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description"`
}
