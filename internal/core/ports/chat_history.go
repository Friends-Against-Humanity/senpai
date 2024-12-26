package ports

type HistoryGateway interface {
	GetHistory(sessionID string) ([]string, error)
}
