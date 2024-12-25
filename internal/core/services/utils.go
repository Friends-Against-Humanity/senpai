package services

import (
	"encoding/json"

	"github.com/Friends-Against-Humanity/senpai/internal/core/domain"
)


func makeMetadata(metadata domain.Metadata) string {
	metadataBytes, _ := json.Marshal(metadata)
	return string(metadataBytes)
}
