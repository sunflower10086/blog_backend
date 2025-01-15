package main

import (
	"sunflower-blog-svc/app/blog/internal/data/gormgen/model"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var (
	g  *gen.Generator
	db *gorm.DB
)
var dsn = "host=127.0.0.1 user=sunflower password=lz18738377974 dbname=blog port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func init() {
	// db, _ = gorm.Open(mysql.Open("root:root@tcp(192.168.127.128:3306)/austin-v2?parseTime=true&collation=utf8mb4_unicode_ci&loc=Asia%2FShanghai&charset=utf8mb4"), &gorm.Config{})

	var err error
	postgresConf := postgres.Config{
		DSN: dsn,
	}
	// gormConfig := configLog(c.Postgres.LogMode, c.Postgres.CreateBatchSize)
	if db, err = gorm.Open(postgres.New(postgresConf)); err != nil {
		log.Fatal("opens database failed: ", err)
	}
}

func main() {
	g = gen.NewGenerator(gen.Config{
		OutPath:           "./query",
		Mode:              gen.WithoutContext | gen.WithQueryInterface,
		FieldNullable:     true,
		FieldWithIndexTag: true,
	})

	g.UseDB(db)
	db.AutoMigrate(&model.Post{})

	// tableList = relationship(tableList) //需要处理关系的表

	// 其他默认的表
	g.ApplyBasic(model.Post{})
	// g.ApplyInterface(func(CommonDao) {}, g.GenerateModel("la_user"))
	g.Execute()
}
