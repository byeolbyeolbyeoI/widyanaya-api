package repository

import (
	"github.com/byeolbyeolbyeoI/widyanaya-api/helper"
	"github.com/byeolbyeolbyeoI/widyanaya-api/internal/publication/model"
	supa "github.com/nedpals/supabase-go"
	"strconv"
	"strings"
	"time"
)

// errnotfound buat lebih spesifik
type PublicationRepository struct {
	client *supa.Client
	helper helper.HelperInstance
}

func NewPublicationRepository(s *supa.Client, h helper.HelperInstance) PublicationRepositoryInstance {
	return &PublicationRepository{
		client: s,
		helper: h,
	}
}

// refactor struct declare tiap handler
/*
func (u *PublicationRepository) GetPublications() ([]model.Publication, error) {
	var publications []model.Publication
	err := u.client.DB.From("publications").Select(
		"id",
		"title",
		"cover_url",
		"description",
		"volume",
		"year",
		"opening_date",
		"closing_date",
		"review_estimation",
		"publisher_id",
		"publication_category_id").Execute(&publications)
	if err != nil {
		return nil, err
	}

	return publications, nil
}
*/

func (u *PublicationRepository) IsPublisherExists(id int) error {
	strId := strconv.Itoa(id)
	var username map[string]interface{}
	err := u.client.DB.From("users").Select("username").Single().Eq("id", strId).Execute(&username)
	if err != nil {
		if strings.Contains(err.Error(), "JSON object requested, multiple (or no) rows returned") {
			return helper.ErrPublisherNotFound
		}
		return err
	}

	return nil
}

type UnparsedPublication struct {
	Id                    int    `json:"id" validate:"required"`
	Title                 string `json:"title" validate:"required,max=255"`
	CoverURL              string `json:"cover_url" validate:"required,url,max=255"`
	Description           string `json:"description"`
	Volume                int    `json:"volume" validate:"min=0"`
	Year                  int    `json:"year" validate:"required,min=1000,max=9999"`
	OpeningDate           string `json:"opening_date"`
	ClosingDate           string `json:"closing_date" validate:"gtefield=OpeningDate"`
	ReviewEstimation      int    `json:"review_estimation"`
	PublisherId           int    `json:"publisher_id"`
	PublicationCategoryID int    `json:"publication_category_id"`
}

func parsePublication(unparsedPublication UnparsedPublication) (model.Publication, error) {
	var publication model.Publication
	openingDate, err := time.Parse("2006-01-02", unparsedPublication.OpeningDate)
	if err != nil {
		return model.Publication{}, err
	}
	closingDate, err := time.Parse("2006-01-02", unparsedPublication.ClosingDate)
	if err != nil {
		return model.Publication{}, err
	}
	publication = model.Publication{
		Id:                    unparsedPublication.Id,
		Title:                 unparsedPublication.Title,
		CoverURL:              unparsedPublication.CoverURL,
		Description:           unparsedPublication.Description,
		Volume:                unparsedPublication.Volume,
		Year:                  unparsedPublication.Year,
		OpeningDate:           openingDate,
		ClosingDate:           closingDate,
		ReviewEstimation:      unparsedPublication.ReviewEstimation,
		PublisherId:           unparsedPublication.PublisherId,
		PublicationCategoryID: unparsedPublication.PublicationCategoryID,
	}

	return publication, nil
}

func (u *PublicationRepository) parsePublications(unparsedPublications []UnparsedPublication) ([]model.Publication, error) {
	var publications []model.Publication
	for _, pub := range unparsedPublications {
		openingDate, err := time.Parse("2006-01-02", pub.OpeningDate)
		if err != nil {
			return nil, err
		}
		closingDate, err := time.Parse("2006-01-02", pub.ClosingDate)
		if err != nil {
			return nil, err
		}
		publications = append(publications, model.Publication{
			Id:                    pub.Id,
			Title:                 pub.Title,
			CoverURL:              pub.CoverURL,
			Description:           pub.Description,
			Volume:                pub.Volume,
			Year:                  pub.Year,
			OpeningDate:           openingDate,
			ClosingDate:           closingDate,
			ReviewEstimation:      pub.ReviewEstimation,
			PublisherId:           pub.PublisherId,
			PublicationCategoryID: pub.PublicationCategoryID,
		})
	}

	return publications, nil
}

// ass code i know ion have the mood for this stupid shit brah im done w life alr
func (u *PublicationRepository) GetPublications() ([]model.Publication, error) {
	var unparsedPublications []UnparsedPublication
	err := u.client.DB.From("publications").Select("*").Execute(&unparsedPublications)
	if err != nil {
		return nil, err
	}

	publications, err := u.parsePublications(unparsedPublications)
	if err != nil {
		return nil, err
	}

	return publications, nil
}

func (u *PublicationRepository) GetPublicationById(id int) (model.Publication, error) {
	stringId := strconv.Itoa(id)
	var unparsedPublications UnparsedPublication
	err := u.client.DB.From("publications").Select("*").Single().Eq("id", stringId).Execute(&unparsedPublications)
	if err != nil {
		return model.Publication{}, err
	}

	publication, err := parsePublication(unparsedPublications)
	if err != nil {
		return model.Publication{}, err
	}

	return publication, nil
}

func (u *PublicationRepository) GetPublicationByVolume() error {

	return nil
}

func (u *PublicationRepository) GetPublicationByYear() error {

	return nil
}

func (u *PublicationRepository) GetPublicationsByCategoryId(categoryId int) ([]model.Publication, error) {
	strCategoryId := strconv.Itoa(categoryId)
	var unparsedPublications []UnparsedPublication
	err := u.client.DB.From("publications").Select("*").Eq("publication_category_id", strCategoryId).Execute(&unparsedPublications)
	if err != nil {
		return nil, err
	}

	publications, err := u.parsePublications(unparsedPublications)
	if err != nil {
		return nil, err
	}

	return publications, nil
}

func (u *PublicationRepository) AddPublication(publication model.Publication) error {
	var empty []map[string]interface{}

	err := u.client.DB.From("publications").Insert(map[string]interface{}{
		"title":                   publication.Title,
		"cover_url":               publication.CoverURL,
		"description":             publication.Description,
		"volume":                  publication.Volume,
		"year":                    publication.Year,
		"opening_date":            publication.OpeningDate,
		"closing_date":            publication.ClosingDate,
		"review_estimation":       publication.ReviewEstimation,
		"publisher_id":            publication.PublisherId,
		"publication_category_id": publication.PublicationCategoryID,
	}).Execute(&empty)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationRepository) UpdatePublication(publication model.UpdatedPublication) error {
	var empty []map[string]interface{}
	strId := strconv.Itoa(publication.Id)

	err := u.client.DB.From("publications").Update(map[string]interface{}{
		"title":                   publication.Title,
		"cover_url":               publication.CoverURL,
		"description":             publication.Description,
		"volume":                  publication.Volume,
		"year":                    publication.Year,
		"opening_date":            publication.OpeningDate,
		"closing_date":            publication.ClosingDate,
		"review_estimation":       publication.ReviewEstimation,
		"publisher_id":            publication.PublisherId,
		"publication_category_id": publication.PublicationCategoryID,
	}).Eq("id", strId).Execute(&empty)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationRepository) IsPublicationExists(id int) error {
	strId := strconv.Itoa(id)
	var title map[string]interface{}
	err := u.client.DB.From("publications").Select("title").Single().Eq("id", strId).Execute(&title)
	if err != nil {
		if strings.Contains(err.Error(), "JSON object requested, multiple (or no) rows returned") {
			return helper.ErrPublicationNotFound
		}
		return err
	}

	return nil
}

func (u *PublicationRepository) IsPublicationsExist() error {
	var res []map[string]interface{}
	err := u.client.DB.From("publications").Select("id").Execute(&res)
	if err != nil {
		return err
	}

	if len(res) == 0 {
		return helper.ErrPublicationNotFound
	}

	return nil
}

func (u *PublicationRepository) IsPublicationCategoryExists(categoryId int) error {
	strCategoryId := strconv.Itoa(categoryId)
	var id map[string]interface{}
	err := u.client.DB.From("publication_categories").Select("id").Single().Eq("id", strCategoryId).Execute(&id)
	if err != nil {
		if strings.Contains(err.Error(), "JSON object requested, multiple (or no) rows returned") {
			return helper.ErrCategoryNotFound
		}
		return err
	}

	return nil
}

func (u *PublicationRepository) IsPublicationsExistByCategoryId(categoryId int) error {
	var res []map[string]interface{}
	strCategoryId := strconv.Itoa(categoryId)
	err := u.client.DB.From("publications").Select("id").Eq("publication_category_id", strCategoryId).Execute(&res)
	if err != nil {
		return err
	}

	if len(res) == 0 {
		return helper.ErrPublicationNotFound
	}

	return nil
}

func (u *PublicationRepository) DeletePublicationById(id int) error {
	strId := strconv.Itoa(id)
	var empty []map[string]interface{}

	err := u.client.DB.From("publications").Delete().Eq("id", strId).Execute(&empty)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationRepository) IsPaperExists(id int) error {
	strId := strconv.Itoa(id)
	var title map[string]interface{}
	err := u.client.DB.From("papers").Select("title").Single().Eq("id", strId).Execute(&title)
	if err != nil {
		if strings.Contains(err.Error(), "JSON object requested, multiple (or no) rows returned") {
			return helper.ErrPaperNotFound
		}
		return err
	}

	return nil
}

func (u *PublicationRepository) IsPapersExist() error {
	var res []map[string]interface{}
	err := u.client.DB.From("papers").Select("id").Execute(&res)
	if err != nil {
		return err
	}

	if len(res) == 0 {
		return helper.ErrPaperNotFound
	}

	return nil
}

func (u *PublicationRepository) GetPapers() ([]model.Paper, error) {
	var papers []model.Paper
	err := u.client.DB.From("papers").Select("*").Execute(&papers)
	if err != nil {
		return nil, err
	}

	return papers, nil
}

func (u *PublicationRepository) GetPaperById(id int) (model.Paper, error) {
	stringId := strconv.Itoa(id)
	var paper model.Paper
	err := u.client.DB.From("papers").Select("*").Single().Eq("id", stringId).Execute(&paper)
	if err != nil {
		return model.Paper{}, err
	}

	return paper, nil
}

func (u *PublicationRepository) AddPaper(paper model.Paper) error {
	var empty []map[string]interface{}

	err := u.client.DB.From("papers").Insert(map[string]interface{}{
		"title":    paper.Title,
		"keywords": paper.Keywords,
		"owner_id": paper.OwnerId,
	}).Execute(&empty)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationRepository) IsOwnerExists(id int) error {
	strId := strconv.Itoa(id)
	var idRes map[string]interface{}
	err := u.client.DB.From("users").Select("id").Single().Eq("id", strId).Execute(&idRes)
	if err != nil {
		if strings.Contains(err.Error(), "JSON object requested, multiple (or no) rows returned") {
			return helper.ErrOwnerNotFound
		}
		return err
	}

	return nil
}

func (u *PublicationRepository) UpdatePaper(paper model.UpdatedPaper) error {
	var empty []map[string]interface{}
	id := strconv.Itoa(paper.Id)

	err := u.client.DB.From("papers").Update(map[string]interface{}{
		"title":    paper.Title,
		"keywords": paper.Keywords,
		"owner_id": paper.OwnerId,
	}).Eq("id", id).Execute(&empty)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationRepository) DeletePaperById(id int) error {
	strId := strconv.Itoa(id)
	var res []map[string]interface{}

	err := u.client.DB.From("papers").Delete().Eq("id", strId).Execute(&res)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationRepository) DeletePaperFragmentById(id int) error {
	strId := strconv.Itoa(id)
	var res []map[string]interface{}

	//competitions
	err := u.client.DB.From("paper_fragments").Delete().Eq("id", strId).Execute(&res)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationRepository) IsCompetitionExists(id int) error {
	strId := strconv.Itoa(id)
	var title map[string]interface{}
	err := u.client.DB.From("competitions").Select("name").Single().Eq("id", strId).Execute(&title)
	if err != nil {
		if strings.Contains(err.Error(), "JSON object requested, multiple (or no) rows returned") {
			return helper.ErrCompetitionNotFound
		}
		return err
	}

	return nil
}

func (u *PublicationRepository) IsCompetitionsExist() error {
	var res []map[string]interface{}
	err := u.client.DB.From("competitions").Select("id").Execute(&res)
	if err != nil {
		return err
	}

	if len(res) == 0 {
		return helper.ErrCompetitionNotFound
	}

	return nil
}

func (u *PublicationRepository) parseCompetitions(unparsedCompetitions []model.UnparsedCompetition) ([]model.Competition, error) {
	var competitions []model.Competition
	for _, pub := range unparsedCompetitions {
		openingDate, err := time.Parse("2006-01-02", pub.OpeningDate)
		if err != nil {
			return nil, err
		}
		closingDate, err := time.Parse("2006-01-02", pub.ClosingDate)
		if err != nil {
			return nil, err
		}

		date, err := time.Parse("2006-01-02", pub.Date)
		if err != nil {
			return nil, err
		}

		competitions = append(competitions, model.Competition{
			Id:                    pub.Id,
			Name:                  pub.Name,
			Description:           pub.Description,
			OpeningDate:           openingDate,
			ClosingDate:           closingDate,
			Date:                  date,
			Fees:                  pub.Fees,
			CompetitionCategoryId: pub.CompetitionCategoryId,
			PublisherId:           pub.PublisherId,
		})
	}

	return competitions, nil
}

func parseCompetition(unparsedCompetition model.UnparsedCompetition) (model.Competition, error) {
	var competition model.Competition
	openingDate, err := time.Parse("2006-01-02", unparsedCompetition.OpeningDate)
	if err != nil {
		return model.Competition{}, err
	}

	closingDate, err := time.Parse("2006-01-02", unparsedCompetition.ClosingDate)
	if err != nil {
		return model.Competition{}, err
	}

	date, err := time.Parse("2006-01-02", unparsedCompetition.Date)
	if err != nil {
		return model.Competition{}, err
	}

	competition = model.Competition{
		Id:                    unparsedCompetition.Id,
		Name:                  unparsedCompetition.Name,
		Description:           unparsedCompetition.Description,
		OpeningDate:           openingDate,
		ClosingDate:           closingDate,
		Date:                  date,
		Fees:                  unparsedCompetition.Fees,
		CompetitionCategoryId: unparsedCompetition.CompetitionCategoryId,
		PublisherId:           unparsedCompetition.PublisherId,
	}

	return competition, nil
}

func (u *PublicationRepository) GetCompetitions() ([]model.Competition, error) {
	var unparsedCompetition []model.UnparsedCompetition
	err := u.client.DB.From("competitions").Select("*").Execute(&unparsedCompetition)
	if err != nil {
		return nil, err
	}

	competitions, err := u.parseCompetitions(unparsedCompetition)
	if err != nil {
		return nil, err
	}

	return competitions, nil
}

func (u *PublicationRepository) GetCompetitionById(id int) (model.Competition, error) {
	stringId := strconv.Itoa(id)
	var unparsedCompetition model.UnparsedCompetition
	err := u.client.DB.From("competitions").Select("*").Single().Eq("id", stringId).Execute(&unparsedCompetition)
	if err != nil {
		return model.Competition{}, err
	}

	competition, err := parseCompetition(unparsedCompetition)
	if err != nil {
		return model.Competition{}, err
	}

	return competition, nil
}

func (u *PublicationRepository) IsCompetitionCategoryExists(categoryId int) error {
	strCategoryId := strconv.Itoa(categoryId)
	var id map[string]interface{}
	err := u.client.DB.From("competition_categories").Select("id").Single().Eq("id", strCategoryId).Execute(&id)
	if err != nil {
		if strings.Contains(err.Error(), "JSON object requested, multiple (or no) rows returned") {
			return helper.ErrCategoryNotFound
		}
		return err
	}

	return nil
}

func (u *PublicationRepository) IsCompetitionsExistByCategoryId(categoryId int) error {
	var res []map[string]interface{}
	strCategoryId := strconv.Itoa(categoryId)
	err := u.client.DB.From("competitions").Select("id").Eq("competition_category_id", strCategoryId).Execute(&res)
	if err != nil {
		return err
	}

	if len(res) == 0 {
		return helper.ErrCompetitionNotFound
	}

	return nil
}

func (u *PublicationRepository) GetCompetitionsByCategoryId(categoryId int) ([]model.Competition, error) {
	strCategoryId := strconv.Itoa(categoryId)
	var unparsedCompetitions []model.UnparsedCompetition
	err := u.client.DB.From("competitions").Select("*").Eq("competition_category_id", strCategoryId).Execute(&unparsedCompetitions)
	if err != nil {
		return nil, err
	}

	competitions, err := u.parseCompetitions(unparsedCompetitions)
	if err != nil {
		return nil, err
	}

	return competitions, nil
}

func (u *PublicationRepository) AddCompetition(competition model.Competition) error {
	var empty []map[string]interface{}

	err := u.client.DB.From("competitions").Insert(map[string]interface{}{
		"name":                    competition.Name,
		"description":             competition.Description,
		"opening_date":            competition.OpeningDate,
		"closing_date":            competition.ClosingDate,
		"date":                    competition.Date,
		"fees":                    competition.Fees,
		"competition_category_id": competition.CompetitionCategoryId,
		"publisher_id":            competition.PublisherId,
	}).Execute(&empty)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationRepository) UpdateCompetition(competition model.UpdatedCompetition) error {
	var empty []map[string]interface{}
	id := strconv.Itoa(competition.Id)
	err := u.client.DB.From("competitions").Update(map[string]interface{}{
		"name":                    competition.Name,
		"description":             competition.Description,
		"opening_date":            competition.OpeningDate,
		"closing_date":            competition.ClosingDate,
		"date":                    competition.Date,
		"fees":                    competition.Fees,
		"competition_category_id": competition.CompetitionCategoryId,
		"publisher_id":            competition.PublisherId,
	}).Eq("id", id).Execute(&empty)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationRepository) DeleteCompetitionById(id int) error {
	strId := strconv.Itoa(id)
	var res []map[string]interface{}

	//competitions
	err := u.client.DB.From("competitions").Delete().Eq("id", strId).Execute(&res)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationRepository) IsPublicationRequestExists(id int) error {
	strId := strconv.Itoa(id)
	var resId map[string]interface{}
	err := u.client.DB.From("publication_requests").Select("id").Single().Eq("id", strId).Execute(&resId)
	if err != nil {
		if strings.Contains(err.Error(), "JSON object requested, multiple (or no) rows returned") {
			return helper.ErrPublicationRequestNotFound
		}
		return err
	}

	return nil
}
func (u *PublicationRepository) DeletePublicationRequestById(id int) error {
	strId := strconv.Itoa(id)
	var res []map[string]interface{}

	//requests?
	err := u.client.DB.From("publication_requests").Delete().Eq("id", strId).Execute(&res)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationRepository) IsPaperFragmentsExist() error {
	var res []map[string]interface{}
	err := u.client.DB.From("paper_fragments").Select("id").Execute(&res)
	if err != nil {
		return err
	}

	if len(res) == 0 {
		return helper.ErrPaperFragmentNotFound
	}

	return nil
}

func (u *PublicationRepository) GetPaperFragments() ([]model.PaperFragment, error) {
	var fragments []model.PaperFragment
	err := u.client.DB.From("paper_fragments").Select("*").Execute(&fragments)
	if err != nil {
		return nil, err
	}

	return fragments, nil
}

func (u *PublicationRepository) GetPaperFragmentById(id int) (model.PaperFragment, error) {
	stringId := strconv.Itoa(id)
	var fragment model.PaperFragment
	err := u.client.DB.From("paper_fragments").Select("*").Single().Eq("id", stringId).Execute(&fragment)
	if err != nil {
		return model.PaperFragment{}, err
	}

	return fragment, nil
}

func (u *PublicationRepository) IsPaperFragmentExists(id int) error {
	strId := strconv.Itoa(id)
	var rowId map[string]interface{}
	err := u.client.DB.From("paper_fragments").Select("id").Single().Eq("id", strId).Execute(&rowId)
	if err != nil {
		if strings.Contains(err.Error(), "JSON object requested, multiple (or no) rows returned") {
			return helper.ErrPaperFragmentNotFound
		}
		return err
	}

	return nil
}

func (u *PublicationRepository) IsPaperFragmentCategoryExists(categoryId int) error {
	strCategoryId := strconv.Itoa(categoryId)
	var id map[string]interface{}
	err := u.client.DB.From("fragment_categories").Select("id").Single().Eq("id", strCategoryId).Execute(&id)
	if err != nil {
		if strings.Contains(err.Error(), "JSON object requested, multiple (or no) rows returned") {
			return helper.ErrCategoryNotFound
		}
		return err
	}

	return nil
}

func (u *PublicationRepository) AddPaperFragment(fragment model.PaperFragment) error {
	var empty []map[string]interface{}

	err := u.client.DB.From("paper_fragments").Insert(map[string]interface{}{
		"content":              fragment.Content,
		"created_at":           fragment.CreatedAt,
		"updated_at":           fragment.UpdatedAt,
		"paper_id":             fragment.PaperId,
		"fragment_category_id": fragment.FragmentCategoryId,
	}).Execute(&empty)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationRepository) UpdatePaperFragment(fragment model.UpdatedPaperFragment) error {
	var empty []map[string]interface{}
	id := strconv.Itoa(fragment.Id)

	err := u.client.DB.From("paper_fragments").Update(map[string]interface{}{
		"content":              fragment.Content,
		"created_at":           fragment.CreatedAt,
		"updated_at":           fragment.UpdatedAt,
		"paper_id":             fragment.PaperId,
		"fragment_category_id": fragment.FragmentCategoryId,
	}).Eq("id", id).Execute(&empty)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationRepository) IsPublicationRequestsExist() error {
	var res []map[string]interface{}
	err := u.client.DB.From("publication_requests").Select("id").Execute(&res)
	if err != nil {
		return err
	}

	if len(res) == 0 {
		return helper.ErrPublicationRequestNotFound
	}

	return nil
}

func (u *PublicationRepository) GetPublicationRequests() ([]model.PublicationRequest, error) {
	var requests []model.PublicationRequest
	err := u.client.DB.From("publication_requests").Select("*").Execute(&requests)
	if err != nil {
		return nil, err
	}

	return requests, nil
}

func (u *PublicationRepository) GetPublicationRequestById(id int) (model.PublicationRequest, error) {
	stringId := strconv.Itoa(id)
	var request model.PublicationRequest
	err := u.client.DB.From("publication_requests").Select("*").Single().Eq("id", stringId).Execute(&request)
	if err != nil {
		return model.PublicationRequest{}, err
	}

	return request, nil
}

func (u *PublicationRepository) AddPublicationRequest(req model.PublicationRequest) error {
	var empty []map[string]interface{}

	err := u.client.DB.From("publication_requests").Insert(map[string]interface{}{
		"paper_url":           req.PaperURL,
		"cover_letter_url":    req.CoverLetterURL,
		"approval_letter_url": req.ApprovalLetterURL,
		"status":              req.Status,
		"metadata_id":         req.MetadataID,
		"reference_format_id": req.ReferenceFormatID,
		"requester_id":        req.RequesterID,
	}).Execute(&empty)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationRepository) UpdatePublicationRequest(req model.UpdatedPublicationRequest) error {
	var empty []map[string]interface{}
	id := strconv.Itoa(req.Id)

	err := u.client.DB.From("publication_requests").Update(map[string]interface{}{
		"paper_url":           req.PaperURL,
		"cover_letter_url":    req.CoverLetterURL,
		"approval_letter_url": req.ApprovalLetterURL,
		"status":              req.Status,
		"metadata_id":         req.MetadataID,
		"reference_format_id": req.ReferenceFormatID,
		"requester_id":        req.RequesterID,
	}).Eq("id", id).Execute(&empty)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationRepository) IsMetadataExists(id int) error {
	strId := strconv.Itoa(id)
	var resId map[string]interface{}
	err := u.client.DB.From("metadata").Select("id").Single().Eq("id", strId).Execute(&resId)
	if err != nil {
		if strings.Contains(err.Error(), "JSON object requested, multiple (or no) rows returned") {
			return helper.ErrMetadataNotFound
		}

		return err
	}

	return nil
}

func (u *PublicationRepository) IsMetadatasExist() error {
	var res []map[string]interface{}
	err := u.client.DB.From("metadata").Select("id").Execute(&res)
	if err != nil {
		return err
	}

	if len(res) == 0 {
		return helper.ErrMetadataNotFound
	}

	return nil
}

func (u *PublicationRepository) IsReferenceFormatExists(formatId int) error {
	strFormatId := strconv.Itoa(formatId)
	var id map[string]interface{}
	err := u.client.DB.From("reference_formats").Select("id").Single().Eq("id", strFormatId).Execute(&id)
	if err != nil {
		if strings.Contains(err.Error(), "JSON object requested, multiple (or no) rows returned") {
			return helper.ErrReferenceFormatNotFound
		}
		return err
	}

	return nil
}

func (u *PublicationRepository) IsRequesterExists(id int) error {
	strId := strconv.Itoa(id)
	var username map[string]interface{}
	err := u.client.DB.From("users").Select("username").Single().Eq("id", strId).Execute(&username)
	if err != nil {
		if strings.Contains(err.Error(), "JSON object requested, multiple (or no) rows returned") {
			return helper.ErrRequesterNotFound
		}
		return err
	}

	return nil
}

func (u *PublicationRepository) GetMetadatas() ([]model.Metadata, error) {
	var metadatas []model.Metadata
	err := u.client.DB.From("metadata").Select("*").Execute(&metadatas)
	if err != nil {
		return nil, err
	}

	return metadatas, nil
}

func (u *PublicationRepository) GetMetadataById(id int) (model.Metadata, error) {
	stringId := strconv.Itoa(id)
	var metadata model.Metadata
	err := u.client.DB.From("metadata").Select("*").Single().Eq("id", stringId).Execute(&metadata)
	if err != nil {
		return model.Metadata{}, err
	}

	return metadata, nil
}

func (u *PublicationRepository) DeleteMetadataById(id int) error {
	strId := strconv.Itoa(id)
	var res []map[string]interface{}

	//requests?
	err := u.client.DB.From("metadata").Delete().Eq("id", strId).Execute(&res)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationRepository) AddMetadata(metadata model.Metadata) error {
	var empty []map[string]interface{}

	err := u.client.DB.From("metadata").Insert(map[string]interface{}{
		"title":          metadata.Title,
		"abstract":       metadata.Abstract,
		"keyword":        metadata.Keyword,
		"contributor":    metadata.Contributor,
		"date_sent":      metadata.DateSent,
		"reference":      metadata.Reference,
		"doi":            metadata.DOI,
		"attachment_url": metadata.AttachmentURL,
	}).Execute(&empty)
	if err != nil {
		return err
	}

	return nil
}

func (u *PublicationRepository) UpdateMetadata(metadata model.UpdatedMetadata) error {
	var empty []map[string]interface{}
	id := strconv.Itoa(metadata.Id)

	err := u.client.DB.From("metadata").Update(map[string]interface{}{
		"title":          metadata.Title,
		"abstract":       metadata.Abstract,
		"keyword":        metadata.Keyword,
		"contributor":    metadata.Contributor,
		"date_sent":      metadata.DateSent,
		"reference":      metadata.Reference,
		"doi":            metadata.DOI,
		"attachment_url": metadata.AttachmentURL,
	}).Eq("id", id).Execute(&empty)
	if err != nil {
		return err
	}

	return nil
}
