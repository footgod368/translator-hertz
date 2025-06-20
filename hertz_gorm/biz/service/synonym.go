package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/hertz_gen/translator"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/llm"
)

const (
	systemPrompt     = "你是一个专业的英语老师，你需要帮用户分析单词的同近义词。"
	userPromptFormat = "请用中文详细分析单词%s的同近义词，每个同近义词包含：\n" +
		"1. 核心含义（不超过15字）\n" +
		"2. 典型例句（中英对照）\n" +
		"3. 与原词的主要差异\n" +
		"格式要求：使用Markdown无序列表和代码块排版\n" +
		"注意排除原词本身\n"
)

func Synonyms(ctx context.Context, req translator.SynonymsRequest) (*translator.SynonymsResponse, error) {
	chatResp, err := llm.Chat(ctx, systemPrompt, fmt.Sprintf(userPromptFormat, req.Word))
	if err != nil {
		return nil, err
	}
	return &translator.SynonymsResponse{
		SynonymsMd: chatResp,
	}, nil
}
