package storage

import (
	"fmt"
	"moges/common/config"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
)

var (
	// dialectMySQL is MySQL dialect
	dialectMySQL = "mysql"
	// dialectSQLite is SQLite dialect
	dialectSQLite = "sqlite3"
	// GormMaster is for master db connection
	GormMaster *gorm.DB
)

func GetDB() *gorm.DB {
	return GormMaster
}

// InitializeGorm ...
func InitializeGorm(cfg *config.MysqlConfig) {
	var err error
	GormMaster, err = gorm.Open(dialectMySQL, cfg.Uri)
	if err != nil {
		panic(fmt.Errorf("failed to init master mysql, %v", err))
	}
	GormMaster.LogMode(cfg.ShowSql)
	GormMaster.DB().SetMaxIdleConns(cfg.MaxIdle)
	GormMaster.DB().SetMaxOpenConns(cfg.MaxOpen)
}

// Close ...
func Close() {
	if GormMaster != nil {
		err := GormMaster.Close()
		if err != nil {

		}
	}
}
