package db2gorm

import "zmusic/music-server/utils/db2gorm/gen"

func generateOne() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/music?charset=utf8&parseTime=true&loc=Local"

	//生成指定单表
	tblName := "consumer"

	gen.GenerateOne(gen.GenConf{
		Dsn:       dsn,
		WritePath: "./model",
		Stdout:    false,
		Overwrite: true,
	}, tblName)
}

func generateAll() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/music?charset=utf8&parseTime=true&loc=Local"

	gen.GenerateAll(gen.GenConf{
		Dsn:       dsn,
		WritePath: "./model",
		Stdout:    false,
		Overwrite: true,
	})
}
