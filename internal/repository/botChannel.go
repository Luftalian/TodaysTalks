package repository

import (
	"context" // Add missing import statement
	"log"
	"os"

	"github.com/traPtitech/go-traq"
)

func (r *Repository2) LetBotJoinChannel(ChannelID string) error {
	_, err := r.apiClient.BotApi.LetBotJoinChannel(context.Background(), os.Getenv("BOT_ID")).
		PostBotActionJoinRequest(*traq.NewPostBotActionJoinRequest(ChannelID)).
		Execute()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *Repository2) LetBotLeaveChannel(ChannelID string) error {
	_, err := r.apiClient.BotApi.LetBotLeaveChannel(context.Background(), os.Getenv("BOT_ID")).
		PostBotActionLeaveRequest(*traq.NewPostBotActionLeaveRequest(ChannelID)).
		Execute()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
