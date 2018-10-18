package mysql

import "github.com/jinzhu/gorm"

// TxFunc is params type of gorm tx
type TxFunc func(tx *gorm.DB) error

// UseTx is transaction wrapper for gorm
func UseTx(db *gorm.DB, funcs... TxFunc) error {
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
