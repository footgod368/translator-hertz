package llm

import (
	"os"

	ark "github.com/sashabaranov/go-openai"
)

var LLMClient *ark.Client
var modelID string = "ep-m-20250611011007-vv5fp"

func Init() {
	config := ark.DefaultConfig(os.Getenv("ARK_API_KEY"))
	config.BaseURL = "https://ark.cn-beijing.volces.com/api/v3"
	LLMClient = ark.NewClientWithConfig(config)
}
