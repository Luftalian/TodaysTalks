package repository

import "context"

func (r *Repository2) AddMessageStamp(MessageID string, StampID string) error {
	_, err := r.apiClient.
		MessageApi.
		AddMessageStamp(context.Background(), MessageID, StampID).
		Execute()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository2) RemoveMessageStamp(MessageID string, StampID string) error {
	_, err := r.apiClient.
		MessageApi.
		RemoveMessageStamp(context.Background(), MessageID, StampID).
		Execute()
	if err != nil {
		return err
	}
	return nil
}
