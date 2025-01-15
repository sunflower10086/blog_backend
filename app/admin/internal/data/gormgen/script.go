package main

import (
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
	"sunflower-blog-svc/app/admin/internal/data/gormgen/model"
)

var g *gen.Generator
var db *gorm.DB
var (
	dsn = "host=127.0.0.1 user=sunflower password=lz18738377974 dbname=user port=5432 sslmode=disable TimeZone=Asia/Shanghai"
)

func init() {
	//db, _ = gorm.Open(mysql.Open("root:root@tcp(192.168.127.128:3306)/austin-v2?parseTime=true&collation=utf8mb4_unicode_ci&loc=Asia%2FShanghai&charset=utf8mb4"), &gorm.Config{})

	var (
		err error
	)
	postgresConf := postgres.Config{
		DSN: dsn,
	}
	//gormConfig := configLog(c.Postgres.LogMode, c.Postgres.CreateBatchSize)
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

	//dataMap := map[string]func(dtype string) string{
	//	//"smallint":  func(dType string) string { return "int32" },
	//	//"tinyint":   func(dType string) string { return "int32" },
	//	//"mediumint": func(dType string) string { return "int32" },
	//	//"bigint":    func(dType string) string { return "int64" },
	//	"timestamptz": func(dType string) string { return "sql.NullTime" },
	//}
	//g.WithDataTypeMap(dataMap)
	//g.FieldNullable

	g.UseDB(db)
	db.AutoMigrate(&model.User{})
	//var tableList []string
	//tableList, _ = db.Migrator().GetTables()
	//
	////tableList = relationship(tableList) //需要处理关系的表
	//
	////其他默认的表
	//for _, v := range tableList {
	//	g.ApplyBasic(g.GenerateModel(v))
	//}

	g.ApplyBasic(model.User{})

	//g.ApplyInterface(func(CommonDao) {}, g.GenerateModel("la_user"))
	g.Execute()
}
