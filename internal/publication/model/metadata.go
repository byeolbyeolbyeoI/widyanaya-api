package model

type Metadata struct {
	Id            int    `json:"id"`
	Title         string `json:"title" validate:"required,max=255"`
	Abstract      string `json:"abstract" validate:"required"`
	Keyword       string `json:"keyword" validate:"required,max=255"`
	Contributor   string `json:"contributor" validate:"required,max=255"`
	DateSent      string `json:"date_sent" validate:"required"`
	Reference     string `json:"reference" validate:"required"`
	DOI           string `json:"doi" validate:"max=100"`
	AttachmentURL string `json:"attachment_url" validate:"max=255"`
}

type UpdatedMetadata struct {
	Id            int    `json:"id" validate:"required"`
	Title         string `json:"title" validate:"required,max=255"`
	Abstract      string `json:"abstract" validate:"required"`
	Keyword       string `json:"keyword" validate:"required,max=255"`
	Contributor   string `json:"contributor" validate:"required,max=255"`
	DateSent      string `json:"date_sent" validate:"required"`
	Reference     string `json:"reference" validate:"required"`
	DOI           string `json:"doi" validate:"max=100"`
	AttachmentURL string `json:"attachment_url" validate:"max=255"`
}
