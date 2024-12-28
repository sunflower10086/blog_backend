package data

import (
	"gorm.io/gorm"
	"sunflower-blog-svc/common/dal/query"
	"sunflower-blog-svc/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewPosterRepo, NewPostgresDB)

// Data .
type Data struct {
	DB *query.Query
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	queryDB := query.Use(db)

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{
		DB: queryDB,
	}, cleanup, nil
}

func NewPostgresDB(conf *conf.Data, logger log.Logger) (*gorm.DB, func(), error) {
	return nil, nil, nil
}
