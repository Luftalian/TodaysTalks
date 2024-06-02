package handler

import (
	"log"

	"github.com/traPtitech/traq-ws-bot/payload"
)

func (h *Handler) OnPingHandler(p *payload.Ping) {
	log.Println("Received PING event: " + p.Base.EventTime.String())
}
