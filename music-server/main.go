package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"zmusic/music-server/config"
	"zmusic/music-server/dao"
)

func main() {
	r := gin.Default()
	//跨域
	r.Use(config.Cors())
	config.InitRouter(r)
	if err := dao.Init(); err != nil {
		panic(err)
	}
	store := cookie.NewStore([]byte("mySecret"))
	// 设置session中间件，参数mySession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mySession", store))
	r.Run(":8888") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
