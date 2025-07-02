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

type SlackActivityRepository interface {
	CreateSlackActivities(ctx context.Context, tx Tx, activity []*domain.SlackActivity) (int64, error)
	UpdateIntegrationToken(ctx context.Context, tx Tx, userID domain.UserID, token *oauth2.Token) error
}

type slackActivityRepository struct {
}

func NewSlackActivityRepository() SlackActivityRepository {
	return &slackActivityRepository{}
}

func (r *slackActivityRepository) CreateSlackActivities(ctx context.Context, tx Tx, activities []*domain.SlackActivity) (int64, error) {
	query := db.New(tx)
	dbActivities := make([]db.CreateSlackActivitiesParams, len(activities))
	for i, activity := range activities {
		dbActivities[i] = db.CreateSlackActivitiesParams{
			UserID:           int64(activity.UserID),
			SlackUserID:      pgtype.Text{String: string(activity.SlackUserID), Valid: activity.SlackUserID != ""},
			Message:          activity.Message,
			Url:              activity.URL,
			ChannelID:        pgtype.Text{String: activity.ChannelID, Valid: activity.ChannelID != ""},
			MessageTimestamp: pgtype.Text{String: activity.MessageTimestamp, Valid: activity.MessageTimestamp != ""},
			CreatedAt:        activity.CreatedAt.Unix(),
		}
	}
	return query.CreateSlackActivities(ctx, dbActivities)
}

func (r *slackActivityRepository) UpdateIntegrationToken(ctx context.Context, tx Tx, userID domain.UserID, token *oauth2.Token) error {
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
	params := db.UpdateSlackIntegrationTokenParams{
		UserID:       int64(userID),
		AccessToken:  []byte(encryptedAccessToken),
		RefreshToken: []byte(encryptedRefreshToken),
		UpdatedAt:    time.Now().Unix(),
	}
	return query.UpdateSlackIntegrationToken(ctx, params)
}
