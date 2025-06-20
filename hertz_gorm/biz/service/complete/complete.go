package complete

import (
	"context"
	"github.com/bytedance/gg/gslice"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/hertz_gen/translator"
	"strings"
)

const maxSuggestions = 10

func Complete(ctx context.Context, req translator.CompleteRequest) (*translator.CompleteResponse, error) {
	return &translator.CompleteResponse{
		Suggestions: gslice.Filter(AllEngWords, func(word string) bool { return strings.HasPrefix(word, req.Prefix) })[:maxSuggestions],
	}, nil
}
