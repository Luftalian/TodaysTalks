package repository

import (
	"context"
)

func (r *Repository2) GetBot(id string, detail bool) (interface{}, error) {
	resp, _, err := r.apiClient.BotApi.GetBot(context.Background(), id).Detail(detail).Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
