package services

import (
	"encoding/json"

	"github.com/Friends-Against-Humanity/senpai/internal/core/domain"
)

func dumpMetadata(metadata domain.Metadata) string {
	metadataBytes, _ := json.Marshal(metadata)
	return string(metadataBytes)
}

func dumpChatHistory(history domain.ChatHistory) string {
	historyBytes, _ := json.Marshal(history)
	return string(historyBytes)
}
