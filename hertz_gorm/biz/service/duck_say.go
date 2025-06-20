package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/hertz_gen/translator"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/utils"
)

func DuckSay(ctx context.Context, req translator.DuckSayRequest) (*translator.DuckSayResponse, error) {
	todayCount, err := CountQueryRecordsToday(ctx)
	if err != nil {
		return nil, err
	}
	text := utils.If(todayCount > 0,
		fmt.Sprintf("嘎嘎，小鸭今天学习了%d个单词", todayCount),
		"嘿嘿，小鸭今天什么也没学",
	)
	return &translator.DuckSayResponse{
		Text: text,
	}, nil
}
