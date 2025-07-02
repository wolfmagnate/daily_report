package infra

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/wolfmagnate/daily_report/collector/domain"
	"github.com/wolfmagnate/daily_report/collector/infra/db"
	"golang.org/x/oauth2"
)

type GithubActivityRepository interface {
	CreateGithubActivities(ctx context.Context, tx Tx, activities []*domain.GithubActivity) (int64, error)
	UpdateLastProcessedEventID(ctx context.Context, tx Tx, userID domain.UserID, eventID string) error
	UpdateGithubIntegrationToken(ctx context.Context, tx Tx, userID domain.UserID, token *oauth2.Token) error
}

type githubActivityRepository struct {
}

func NewGithubActivityRepository() GithubActivityRepository {
	return &githubActivityRepository{}
}

func (r *githubActivityRepository) CreateGithubActivities(ctx context.Context, tx Tx, activities []*domain.GithubActivity) (int64, error) {
	query := db.New(tx)
	dbActivities := make([]db.CreateGithubActivitiesParams, len(activities))
	for i, activity := range activities {
		dbActivities[i] = db.CreateGithubActivitiesParams{
			UserID:         int64(activity.UserID),
			EventID:        activity.EventID,
			Type:           activity.Type,
			Url:            activity.URL,
			RepositoryName: pgtype.Text{String: activity.RepositoryName, Valid: activity.RepositoryName != ""},
			Payload:        activity.Payload,
			CreatedAt:      activity.CreatedAt.Unix(),
		}
	}
	return query.CreateGithubActivities(ctx, dbActivities)
}

func (r *githubActivityRepository) UpdateLastProcessedEventID(ctx context.Context, tx Tx, userID domain.UserID, eventID string) error {
	query := db.New(tx)
	params := db.UpdateGithubIntegrationLastProcessedEventIDParams{
		UserID:               int64(userID),
		LastProcessedEventID: pgtype.Text{String: eventID, Valid: eventID != ""},
		UpdatedAt:            time.Now().Unix(),
	}
	return query.UpdateGithubIntegrationLastProcessedEventID(ctx, params)
}

func (r *githubActivityRepository) UpdateGithubIntegrationToken(ctx context.Context, tx Tx, userID domain.UserID, token *oauth2.Token) error {
	crypter, err := NewCrypter()
	if err != nil {
		return fmt.Errorf("failed to create crypter: %w", err)
	}

	encryptedAccessToken, err := crypter.Encrypt(token.AccessToken)
	if err != nil {
		return fmt.Errorf("failed to encrypt access token: %w", err)
	}

	encryptedRefreshToken, err := crypter.Encrypt(token.RefreshToken)
	if err != nil {
		return fmt.Errorf("failed to encrypt refresh token: %w", err)
	}

	query := db.New(tx)
	params := db.UpdateGithubIntegrationTokenParams{
		UserID:       int64(userID),
		AccessToken:  []byte(encryptedAccessToken),
		RefreshToken: []byte(encryptedRefreshToken),
		UpdatedAt:    time.Now().Unix(),
	}
	return query.UpdateGithubIntegrationToken(ctx, params)
}
