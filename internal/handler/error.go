package handler

import "log"

func (h *Handler) OnErrorHandler(message string) {
	log.Println("Received ERROR message: " + message)
}
