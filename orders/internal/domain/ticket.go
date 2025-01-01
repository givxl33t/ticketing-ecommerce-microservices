package domain

type Ticket struct {
	ID    int32  `gorm:"column:id"`
	Title string `gorm:"column:title"`
	Price int64  `gorm:"column:price"`
}
