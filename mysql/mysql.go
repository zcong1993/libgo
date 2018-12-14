package mysql

import (
	"github.com/jinzhu/gorm"
)

// TxFunc is params type of gorm tx
type TxFunc func(tx *gorm.DB) error

// UseTx is transaction wrapper for gorm
func UseTx(db *gorm.DB, funcs ...TxFunc) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, f := range funcs {
		err := f(tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// WithPagination warp a orm query, return total count and query instance
func WithPagination(db *gorm.DB, limit, offset int) (int, *gorm.DB, error) {
	var count int
	count = 0
	err := db.Count(&count).Error

	if err != nil {
		return count, nil, err
	}

	return count, db.Limit(limit).Offset(offset), nil
}

func PaginationQuery(db *gorm.DB, limit, offset int, t interface{}) (int, error) {
	c, q, err := WithPagination(db, limit, offset)
	if err != nil {
		return c, err
	}

	err = q.Find(&t).Error
	if err != nil {
		return c, err
	}

	return c, nil
}
