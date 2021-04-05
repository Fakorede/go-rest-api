package interfaces

import "gorm.io/gorm"

type Entity interface {
	Count(db *gorm.DB) int64
	Paginate(db *gorm.DB, limit int, offset int) interface{}
}
