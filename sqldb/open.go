package sqldb

import (
	"github.com/go-sql-driver/mysql"
	"gorm.io/driver/gaussdb"
	gormysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Open 连接数据库，并检测是否是信创 OpenGauss。
func Open(dsn string, opts ...gorm.Option) (*gorm.DB, error) {
	// 区分 MySQL 还是 OpenGauss，用的是 parse dsn 的方式。
	// 注意：用该方式区分必须要先用 [mysql.ParseDSN] 才较为稳妥。
	var dia gorm.Dialector
	if cfg, err := mysql.ParseDSN(dsn); err == nil && cfg != nil {
		dia = &gormysql.Dialector{
			Config: &gormysql.Config{DSN: dsn},
		}
	} else {
		dia = gaussdb.New(gaussdb.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		})
	}

	db, err := gorm.Open(dia, opts...)
	if err != nil {
		return nil, err
	}

	return db, nil
}
