package dal

import (
	"gitee.com/fat_marmota/infra/db_client"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabaseClient(masterDSN string, slaveSDNs []string, dbType string) {
	DB = db_client.InitDatabaseClient(masterDSN, slaveSDNs, dbType)
}
