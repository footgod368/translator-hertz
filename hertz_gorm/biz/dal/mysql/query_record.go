package mysql

import (
	"context"

	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/model"
)

func CreateQueryRecord(ctx context.Context, queryRecord *model.QueryRecord) error {
	return DB.WithContext(ctx).Create(queryRecord).Error
}