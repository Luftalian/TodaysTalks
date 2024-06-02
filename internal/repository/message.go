package repository

import (
	"context"
	"time"

	"github.com/traPtitech/go-traq"
)

type (
	MessageLimitationParams struct {
		Limit     int32
		Offset    int32
		Since     time.Time
		Until     time.Time
		Inclusive bool
		Order     string
	}
)

func (r *Repository2) PostMessage(message string, embed bool, ChannelID string) error {
	_, _, err := r.apiClient.
		MessageApi.
		PostMessage(context.Background(), ChannelID).
		PostMessageRequest(traq.PostMessageRequest{
			Content: message,
			Embed:   &embed,
		}).
		Execute()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository2) GetMessage(ChannelID string, params *MessageLimitationParams) ([]traq.Message, error) {
	limit := int32(0)
	offset := int32(0)
	since := time.Time{}
	until := time.Time{}
	inclusive := false
	order := ""

	if params != nil {
		limit = params.Limit
		offset = params.Offset
		since = params.Since
		until = params.Until
		inclusive = params.Inclusive
		order = params.Order
	}

	resp, _, err := r.apiClient.
		MessageApi.
		GetMessages(context.Background(), ChannelID).
		Limit(limit).
		Offset(offset).
		Since(since).
		Until(until).
		Inclusive(inclusive).
		Order(order).
		Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
