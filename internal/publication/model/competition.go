package model

import "time"

type CompetitionCategory struct {
	Id          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description"`
}

type Competition struct {
	Id                    int       `json:"id"`
	Name                  string    `json:"name" validate:"required,max=255"`
	Description           string    `json:"description" validate:"required"`
	OpeningDate           time.Time `json:"opening_date" validate:"required"`
	ClosingDate           time.Time `json:"closing_date" validate:"required"`
	Date                  time.Time `json:"date" validate:"required"`
	Fees                  float64   `json:"fees" validate:"required"`
	CompetitionCategoryId int       `json:"competition_category_id" validate:"required"`
	PublisherId           int       `json:"publisher_id" validate:"required"`
}

type UpdatedCompetition struct {
	Id                    int       `json:"id" validate:"updated"`
	Name                  string    `json:"name" validate:"required,max=255"`
	Description           string    `json:"description" validate:"required"`
	OpeningDate           time.Time `json:"opening_date" validate:"required"`
	ClosingDate           time.Time `json:"closing_date" validate:"required"`
	Date                  time.Time `json:"date" validate:"required"`
	Fees                  float64   `json:"fees" validate:"required"`
	CompetitionCategoryId int       `json:"competition_category_id" validate:"required"`
	PublisherId           int       `json:"publisher_id" validate:"required"`
}

type UnparsedCompetition struct {
	Id                    int     `json:"id"`
	Name                  string  `json:"name" validate:"required,max=255"`
	Description           string  `json:"description" validate:"required"`
	OpeningDate           string  `json:"opening_date" validate:"required"`
	ClosingDate           string  `json:"closing_date" validate:"required"`
	Date                  string  `json:"date" validate:"required"`
	Fees                  float64 `json:"fees" validate:"required"`
	CompetitionCategoryId int     `json:"competition_category_id" validate:"required"`
	PublisherId           int     `json:"publisher_id" validate:"required"`
}
