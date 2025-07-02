package domain

import (
	"fmt"
	"time"
)

type SlackActivityID int64

type SlackActivity struct {
	ID               SlackActivityID `validate:"required,gt=0"`
	UserID           UserID          `validate:"required,gt=0"`
	SlackUserID      SlackUserID     `validate:"required"`
	Message          string          `validate:"required"`
	URL              string          `validate:"required,url"`
	ChannelID        string          `validate:"required"`
	MessageTimestamp string          `validate:"required"`
	CreatedAt        time.Time       `validate:"required"`
}

func NewSlackActivity(
	id int64,
	userID UserID,
	slackUserID SlackUserID,
	message, url, channelID, messageTimestamp string,
	createdAt int64,
) (*SlackActivity, error) {
	activity := &SlackActivity{
		ID:               SlackActivityID(id),
		UserID:           userID,
		SlackUserID:      slackUserID,
		Message:          message,
		URL:              url,
		ChannelID:        channelID,
		MessageTimestamp: messageTimestamp,
		CreatedAt:        time.Unix(createdAt, 0),
	}

	err := validate.Struct(activity)
	if err != nil {
		return nil, fmt.Errorf("validation failed for slack activity: %w", err)
	}

	return activity, nil
}
