package model

import "time"

type QueryRecord struct {
    ID        uint      `gorm:"primaryKey;autoIncrement"`
    Timestamp time.Time `gorm:"type:datetime(3);default:CURRENT_TIMESTAMP(3)"`
    Word      string    `gorm:"type:text;not null"`
    IPAddress string    `gorm:"type:text"`
    UserAgent string    `gorm:"type:text"`
}

func (QueryRecord) TableName() string {
    return "query_records"
}