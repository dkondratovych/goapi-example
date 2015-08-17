package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/seesawlabs/Dima-Kondravotych-Exercise/shared/config"
)

const (
	dialect = "mysql"
)

func NewGormDb(c *config.Database) (*gorm.DB, error) {
	// Gorm do database ping so here we can be sure that connection is alive
	db, err := gorm.Open(dialect, c.Dsn())

	if err != nil {
		return nil, err
	}

	db.SingularTable(c.GormSingularTable)
	db.LogMode(c.GormLogMode)

	return &db, nil
}
