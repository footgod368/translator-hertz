package service

import (
	"context"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/hertz_gen/translator"
)

func Complete(ctx context.Context, req translator.CompleteRequest) (*translator.CompleteResponse, error) {
	return &translator.CompleteResponse{
		Suggestions: []string{req.Prefix + "111", req.Prefix + "222"},
	}, nil
}
