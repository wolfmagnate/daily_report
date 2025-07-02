package domain

import (
	"fmt"
	"time"

	"golang.org/x/oauth2"
)

type UserID int64

type GithubUserID int32

type SlackUserID string

type GithubIntegration struct {
	Token                *oauth2.Token
	Scopes               string
	GithubUsername       string
	GithubUserID         GithubUserID
	LastProcessedEventID string
}

type SlackIntegration struct {
	Token                *oauth2.Token
	RefreshToken         string
	Scopes               string
	SlackUserID          SlackUserID
	IntegratedChannelIDs []string
	TeamID               string
}

type User struct {
	ID                UserID `validate:"required,gt=0"`
	LastCollectedAt   time.Time
	GithubIntegration *GithubIntegration
	SlackIntegration  *SlackIntegration
}

func NewUser(
	id int64,
	lastCollectedAt int64,
	// GitHub-related parameters
	githubAccessToken string,
	githubRefreshToken string,
	githubScopes string,
	githubUserID int32,
	githubUsername string,
	githubLastProcessedEventID string,
	// Slack-related parameters
	slackAccessToken string,
	slackRefreshToken string,
	slackScopes string,
	slackUserID string,
	slackTeamID string,
	slackChannelIDs []string,
) (*User, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, fmt.Errorf("failed to load timezone 'Asia/Tokyo': %w", err)
	}

	user := &User{
		ID:              UserID(id),
		LastCollectedAt: time.Unix(lastCollectedAt, 0).In(jst),
	}

	if githubAccessToken != "" {
		user.GithubIntegration = &GithubIntegration{
			Token: &oauth2.Token{
				AccessToken:  githubAccessToken,
				RefreshToken: githubRefreshToken,
				TokenType:    "Bearer",
			},
			Scopes:               githubScopes,
			GithubUserID:         GithubUserID(githubUserID),
			GithubUsername:       githubUsername,
			LastProcessedEventID: githubLastProcessedEventID,
		}
	}

	if slackAccessToken != "" {
		user.SlackIntegration = &SlackIntegration{
			Token: &oauth2.Token{
				AccessToken:  slackAccessToken,
				RefreshToken: slackRefreshToken,
				TokenType:    "Bearer",
			},
			Scopes:               slackScopes,
			SlackUserID:          SlackUserID(slackUserID),
			TeamID:               slackTeamID,
			IntegratedChannelIDs: slackChannelIDs,
		}
	}

	err = validate.Struct(user)
	if err != nil {
		return nil, fmt.Errorf("validation failed for user: %w", err)
	}

	return user, nil
}
