package replicate

import "one-api/dto"

type ReplicateRequest struct {
	FrequencyPenalty uint    `json:"frequency_penalty"`
	MaxTokens        uint    `json:"max_tokens"`
	MinTokens        uint    `json:"min_tokens"`
	PresencePenalty  float64 `json:"presence_penalty"`
	Prompt           string  `json:"prompt"`
	SystemPrompt     string  `json:"system_prompt"`
	Temperature      float64 `json:"temperature"`
	TopK             uint    `json:"top_k"`
	TopP             float64 `json:"top_p"`
}

func requestOpenAI2Replicate(request dto.GeneralOpenAIRequest) *ReplicateRequest {
	replicateRequest := ReplicateRequest{
		Temperature:     request.Temperature,
		TopP:            request.TopP,
		MaxTokens:       request.MaxTokens,
		MinTokens:       0,
		PresencePenalty: request.PresencePenalty,
		SystemPrompt:    "",
		TopK:            uint(request.TopK),
	}
	if request.MaxTokens != 0 {
		maxTokens := uint(request.MaxTokens)
		if request.MaxTokens == 1 {
			maxTokens = 2
		}
		replicateRequest.MaxTokens = maxTokens
	}
	for _, message := range request.Messages {
		if message.Role == "system" {
			replicateRequest.SystemPrompt = message.StringContent()
		} else {
			replicateRequest.Prompt = replicateRequest.Prompt + string(message.Content)
		}
	}
	return &replicateRequest
}
