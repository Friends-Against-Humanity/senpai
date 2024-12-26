package handler

import (
	"github.com/Friends-Against-Humanity/senpai/internal/core/services"
	"github.com/Friends-Against-Humanity/senpai/pkg/channels/discord"
)

type Handler struct {
	Bot     *discord.Bot
	Service services.MainService
}

type HandlerConfigurer func(b *Handler)

func NewHandler(cfg ...HandlerConfigurer) (*Handler, error) {
	h := &Handler{}

	for _, c := range cfg {
		c(h)
	}

	h.Bot.Handler(h.entrypoint)

	return h, nil
}
