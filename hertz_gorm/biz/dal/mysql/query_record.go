package mysql

import (
	"context"

	"github.com/cloudwego/hertz-examples/bizdemo/hertz_gorm/biz/model"
)

func CreateQueryRecord(ctx context.Context, queryRecord *model.QueryRecord) error {
	return DB.WithContext(ctx).Create(queryRecord).Error
}

func CountQueryRecords(ctx context.Context, whereClause string, args ...interface{}) (int64, error) {
	var count int64
	err := DB.WithContext(ctx).Model(&model.QueryRecord{}).Where(whereClause, args...).Count(&count).Error
	return count, err
}
