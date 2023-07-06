package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"helloworld/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewDB, NewUserRepo)

// Data .
type Data struct {
	DB *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{DB: db}, cleanup, nil
}

// NewDB .
func NewDB(c *conf.Data) *gorm.DB {

	db, err := gorm.Open(mysql.Open("admin:zwn123456@tcp(43.143.227.115:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic(err)
	}
	return db
}
