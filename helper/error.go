package helper

import "errors"

var ErrNotFound = errors.New("resource not found")

// js realized that i js need to have 1 error not found
var ErrPaperNotFound = errors.New("paper not found")
var ErrCategoryNotFound = errors.New("category not found")
var ErrUserNotFound = errors.New("user not found")
var ErrPublicationNotFound = errors.New("publication not found")
var ErrPaperFragmentNotFound = errors.New("paper fragment not found")
var ErrPublicationRequestNotFound = errors.New("publication request not found")
var ErrCompetitionNotFound = errors.New("competition not found")

var ErrPublisherNotFound = errors.New("publisher not found")
var ErrOwnerNotFound = errors.New("owner not found")
var ErrReferenceFormatNotFound = errors.New("reference format not found")
var ErrRequesterNotFound = errors.New("requester not found")
var ErrMetadataNotFound = errors.New("metadata not found")
