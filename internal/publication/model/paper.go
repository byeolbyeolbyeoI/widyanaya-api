package model

import (
	"time"
)

type FragmentCategory struct {
	Id          int    `json:"id"`
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description"`
}

type PaperFragment struct {
	Id                 int       `json:"id"`
	Content            string    `json:"content" validate:"required"`
	CreatedAt          time.Time `json:"created_at" validate:"required"`
	UpdatedAt          time.Time `json:"updated_at"`
	FragmentCategoryId int       `json:"fragment_category_id" validate:"required"`
	PaperId            int       `json:"paper_id" validate:"required"`
}

type UpdatedPaperFragment struct {
	Id                 int       `json:"id" validate:"required"`
	Content            string    `json:"content" validate:"required"`
	CreatedAt          time.Time `json:"created_at" validate:"required"`
	UpdatedAt          time.Time `json:"updated_at"`
	FragmentCategoryId int       `json:"fragment_category_id" validate:"required"`
	PaperId            int       `json:"paper_id" validate:"required"`
}

type Paper struct {
	Id       int    `json:"id"`
	Title    string `json:"title" validate:"required,max=255"`
	Keywords string `json:"keywords" validate:"required,max=255"`
	OwnerId  int    `json:"owner_id" validate:"required" validate:"required"`
}

type UpdatedPaper struct {
	Id       int    `json:"id" validate:"required"`
	Title    string `json:"title" validate:"required,max=255"`
	Keywords string `json:"keywords" validate:"required,max=255"`
	OwnerId  int    `json:"owner_id" validate:"required" validate:"required"`
}
