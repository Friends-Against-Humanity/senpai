package openai

import "encoding/json"

type Choice struct {
	Message Message `json:"message"`
}

type Response struct {
	Choices []Choice `json:"choices"`
}

func NewResponse(body []byte) (Response, error) {
	response := Response{}
	err := json.Unmarshal(body, &response)
	if err != nil {
		return Response{}, err
	}

	return response, nil
}
