package handler

import (
	"log"
	"os"

	"github.com/traPtitech/traq-ws-bot/payload"
)

func (h *Handler) OnBotMessageStampsUpdatedHandler(p *payload.BotMessageStampsUpdated) {
	log.Println("Received BOT_MESSAGE_STAMPS_UPDATED event: " + p.MessageID)
	for _, stamp := range p.Stamps {
		log.Println(stamp.StampID)
		if stamp.UserID == os.Getenv("USER_ME_ID") && stamp.StampID == "" {
			log.Println("User me stamped")
		}
	}
}
