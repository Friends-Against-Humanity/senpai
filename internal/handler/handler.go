package handler

import (
	"github.com/Friends-Against-Humanity/senpai/internal/core/services"
	"github.com/Friends-Against-Humanity/senpai/pkg/channels/discord"
)

type Handler struct {
	Bot         *discord.Bot
	Service     services.Service
	PersonasMap map[string]string
}

type HandlerConfigurer func(b *Handler)

func NewHandler(cfg ...HandlerConfigurer) (*Handler, error) {
	h := &Handler{
		PersonasMap: map[string]string{},
	}

	for _, c := range cfg {
		c(h)
	}

	h.Bot.Handler(h.entrypoint)

	return h, nil
}
