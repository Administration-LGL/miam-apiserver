package data

import (
	"miam-apiserver/internal/biz"
	"miam-apiserver/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewGormDB,
	NewPolicyRepo,
	NewUserRepo,
	NewSecretRepo,
)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{
		db: db,
	}, cleanup, nil
}

func NewGormDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "miam-apiserver/data/gorm"))
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	db = db.Debug()
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	// 自动迁移数据库表
	if err := db.AutoMigrate(&biz.User{}, &biz.Policy{}, &biz.Secret{}); err != nil {
		log.Fatal(err)
	}
	return db
}
