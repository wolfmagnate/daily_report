package infra

import (
	"context"
	"fmt"

	"github.com/wolfmagnate/daily_report/collector/domain"
	"github.com/wolfmagnate/daily_report/collector/infra/db"
)

type UserRepository interface {
	FindCandidates(ctx context.Context, tx Tx, limit int, excludeIDs []int64) ([]*domain.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindCandidates(ctx context.Context, tx Tx, limit int, excludeIDs []int64) ([]*domain.User, error) {
	crypter, err := NewCrypter()
	if err != nil {
		return nil, fmt.Errorf("failed to create crypter: %w", err)
	}

	query := db.New(tx)
	params := db.FindCandidatesParams{
		Lim:        int32(limit),
		ExcludeIds: excludeIDs,
	}

	rows, err := query.FindCandidates(ctx, params)
	if err != nil {
		return nil, err
	}

	users := make([]*domain.User, len(rows))
	for i, row := range rows {
		var githubAccessToken, githubRefreshToken, githubScopes, githubLastProcessedEventID string
		decrypted, err := crypter.Decrypt(string(row.GiAccessToken))
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt github access token: %w", err)
		}
		githubAccessToken = decrypted
		decrypted, err = crypter.Decrypt(string(row.GiRefreshToken))
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt github refresh token: %w", err)
		}
		githubRefreshToken = decrypted
		if row.GiScopes.Valid {
			githubScopes = row.GiScopes.String
		}
		if row.GiLastProcessedEventID.Valid {
			githubLastProcessedEventID = row.GiLastProcessedEventID.String
		}

		var slackAccessToken, slackRefreshToken, slackScopes, slackUserID, slackTeamID string
		var slackChannelIDs []string
		decrypted, err = crypter.Decrypt(string(row.SiAccessToken))
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt slack access token: %w", err)
		}
		slackAccessToken = decrypted

		decrypted, err = crypter.Decrypt(string(row.SiRefreshToken))
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt slack refresh token: %w", err)
		}
		slackRefreshToken = decrypted

		if row.SiScopes.Valid {
			slackScopes = row.SiScopes.String
		}
		if row.SiSlackUserID.Valid {
			slackUserID = row.SiSlackUserID.String
		}
		if row.SiTeamID.Valid {
			slackTeamID = row.SiTeamID.String
		}
		if row.SiChannelIds != nil {
			slackChannelIDs = row.SiChannelIds.([]string)
		}

		user, err := domain.NewUser(
			row.UID,
			row.ULastCollectedAt,
			githubAccessToken,
			githubRefreshToken,
			githubScopes,
			row.UGithubID,
			row.UGithubUsername,
			githubLastProcessedEventID,
			slackAccessToken,
			slackRefreshToken,
			slackScopes,
			slackUserID,
			slackTeamID,
			slackChannelIDs,
		)
		if err != nil {
			return nil, err // Handle error from NewUser
		}
		users[i] = user
	}

	return users, nil
}
