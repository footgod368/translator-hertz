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

func CountQueryRecordsToday(ctx context.Context) (int64, error) {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return 0, err
	}
	today := time.Now().In(location).Format("2006-01-02")
	return mysql.CountQueryRecords(ctx, "date(timestamp) = ?", today)
}
