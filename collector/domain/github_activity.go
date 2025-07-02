package domain

import (
	"fmt"
	"time"
)

type GithubActivityID int64

type GithubActivity struct {
	ID             GithubActivityID `validate:"required,gt=0"`
	UserID         UserID           `validate:"required,gt=0"`
	EventID        string           `validate:"required"`
	Type           string           `validate:"required"`
	URL            string
	RepositoryName string
	Payload        []byte    `validate:"required"`
	CreatedAt      time.Time `validate:"required"`
}

func NewGithubActivity(
	id int64,
	userID UserID,
	eventID string,
	typeStr string,
	url string,
	repositoryName string,
	payload []byte,
	createdAt int64,
) (*GithubActivity, error) {
	activity := &GithubActivity{
		ID:             GithubActivityID(id),
		UserID:         userID,
		EventID:        eventID,
		Type:           typeStr,
		URL:            url,
		RepositoryName: repositoryName,
		Payload:        payload,
		CreatedAt:      time.Unix(createdAt, 0),
	}

	err := validate.Struct(activity)
	if err != nil {
		return nil, fmt.Errorf("validation failed for github activity: %w", err)
	}

	return activity, nil
}
