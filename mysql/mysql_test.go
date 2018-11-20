package mysql_test

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/libgo/mysql"
	"os"
	"testing"
)

const dbFile = "./test.db"

type User struct {
	Name string
}

func setUpDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", dbFile)
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})

	return db
}

func tearDown() {
	os.Remove(dbFile)
}

func TestUseTx(t *testing.T) {
	db := setUpDB()
	err := mysql.UseTx(db, func(db *gorm.DB) error {
		err := db.Save(&User{"name1"}).Error
		err = db.Save(&User{"name2"}).Error
		err = db.Save(&User{"name3"}).Error
		if err != nil {
			return err
		}
		return errors.New("xsxs")
	})

	assert.Error(t, err)
	var count int
	err = db.Model(new(User)).Count(&count).Error
	assert.Nil(t, err)
	assert.Equal(t, 0, count)

	tearDown()
}
