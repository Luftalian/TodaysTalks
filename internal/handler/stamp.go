package handler

import (
	"log"
	"os"

	"github.com/traPtitech/traq-ws-bot/payload"
)

var (
	StampDoneID = "aea52f9a-7484-47ed-ab8f-3b4cc84a474d"
)

func (h *Handler) OnBotMessageStampsUpdatedHandler(p *payload.BotMessageStampsUpdated) {
	log.Println("Received BOT_MESSAGE_STAMPS_UPDATED event: " + p.MessageID)
	for _, stamp := range p.Stamps {
		log.Println(stamp.StampID)
		if stamp.UserID == os.Getenv("USER_ME_ID") && stamp.StampID == StampDoneID {
			err := h.repo2.DeleteMessage(p.MessageID)
			if err != nil {
				log.Println("Failed to delete message: " + err.Error())
				return
			}
		}
	}
}
