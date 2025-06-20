package service

import (
	"context"

	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/hertz_gen/translator"
)

func Translate(ctx context.Context, req translator.TranslateRequest) (translator.TranslateResponse, error) {
	return translator.TranslateResponse{
		Translation: req.Text + "翻译",
	}, nil
}
