package model

type PublicationRequest struct {
	Id                int    `json:"id"`
	PaperURL          string `json:"paper_url" validate:"required,max=255,url"`
	CoverLetterURL    string `json:"cover_letter_url" validate:"required,max=255,url"`
	ApprovalLetterURL string `json:"approval_letter_url" validate:"required,max=255,url"`
	Status            string `json:"status" validate:"required,oneof='Pending' 'Under Review' 'Approved' 'Rejected"`
	MetadataID        int    `json:"metadata_id" validate:"required"`
	ReferenceFormatID int    `json:"reference_format_id" validate:"required"`
	RequesterID       int    `json:"requester_id" validate:"required"`
}

type UpdatedPublicationRequest struct {
	Id                int    `json:"id"`
	PaperURL          string `json:"paper_url" validate:"required,max=255,url"`
	CoverLetterURL    string `json:"cover_letter_url" validate:"required,max=255,url"`
	ApprovalLetterURL string `json:"approval_letter_url" validate:"required,max=255,url"`
	Status            string `json:"status" validate:"required,oneof='Pending' 'Under Review' 'Approved' 'Rejected"`
	MetadataID        int    `json:"metadata_id" validate:"required"`
	ReferenceFormatID int    `json:"reference_format_id" validate:"required"`
	RequesterID       int    `json:"requester_id" validate:"required"`
}
