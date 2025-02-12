package domain

import "database/sql"

type Ticket struct {
	ID        int32         `gorm:"primaryKey;autoIncrement;column:id"`
	Title     string        `gorm:"column:title"`
	Price     int64         `gorm:"column:price"`
	UserID    string        `gorm:"column:user_id"`
	OrderID   sql.NullInt32 `gorm:"column:order_id;default:null"` // do this for NULLABLE values!!
	CreatedAt int64         `gorm:"column:created_at;autoUpdateTime:milli"`
	UpdatedAt int64         `gorm:"column:updated_at;autoUpdateTime:milli"`
}
