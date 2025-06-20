package service

import (
	"context"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/hertz_gen/translator"
)

func Synonyms(ctx context.Context, req translator.SynonymsRequest) (*translator.SynonymsResponse, error) {
	return &translator.SynonymsResponse{
		SynonymsMd: "## " + req.Word + "\n\n暂无同义词",
	}, nil
}
