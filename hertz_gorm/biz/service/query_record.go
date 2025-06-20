package service

import (
	"context"
	"time"

	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/dal/mysql"
	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/model"
)

func CreateQueryRecord(ctx context.Context, word string, ipAddress string, userAgent string) error {
	queryRecord := &model.QueryRecord{
		Timestamp: time.Now(),
		Word:      word,
		IPAddress: ipAddress,
		UserAgent: userAgent,
	}
	return mysql.CreateQueryRecord(ctx, queryRecord)
}
