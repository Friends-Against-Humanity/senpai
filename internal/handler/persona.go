package handler

func (h *Handler) getPersona(channelID string) string {
	if _, ok := h.PersonasMap[channelID]; !ok {
		h.PersonasMap[channelID] = h.Service.FallbackPersona.GetPersonaName()
	}

	return h.PersonasMap[channelID]
}

func (h *Handler) setPersona(channelID string, persona string) {
	h.PersonasMap[channelID] = persona
}
