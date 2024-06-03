package repository

import (
	"context"
	"fmt"
)

type (
	// channels table
	Channel struct {
		ID     string `db:"id"`
		Name   string `db:"name"`
		UserID string `db:"user_id"`
	}

	SubscribeChannelParams struct {
		ID     string
		Name   string
		UserID string
	}
)

func (r *Repository) GetChannels(ctx context.Context) ([]*Channel, error) {
	channels := []*Channel{}
	if err := r.db.SelectContext(ctx, &channels, "SELECT * FROM channels"); err != nil {
		return nil, fmt.Errorf("select channels: %w", err)
	}

	return channels, nil
}

func (r *Repository) SubscribeChannel(ctx context.Context, params SubscribeChannelParams) error {
	if _, err := r.db.ExecContext(ctx, "INSERT INTO channels (id, name) VALUES (?, ?, ?)", params.ID, params.Name, params.UserID); err != nil {
		return fmt.Errorf("insert channel: %w", err)
	}

	return nil
}

func (r *Repository) UnsubscribeChannel(ctx context.Context, channelID string) error {
	if _, err := r.db.ExecContext(ctx, "DELETE FROM channels WHERE id = ?", channelID); err != nil {
		return fmt.Errorf("delete channel: %w", err)
	}

	return nil
}

func (r *Repository) GetChannelByID(ctx context.Context, channelID string) (*Channel, error) {
	channel := &Channel{}
	if err := r.db.GetContext(ctx, channel, "SELECT * FROM channels WHERE id = ?", channelID); err != nil {
		return nil, fmt.Errorf("select	channel: %w", err)
	}

	return channel, nil
}
