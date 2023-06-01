package shorturl

import (
	"github.com/fimreal/goutils/ezap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AddRoute(g *gin.Engine) {
	if !viper.GetBool(ID) && !viper.GetBool("allservices") {
		return
	}
	err := initDB()
	if err != nil {
		ezap.Error(err.Error())
		ezap.Error("[module] shorturl init faild. Skip loading module")
		return
	}

	// help info
	g.GET("/help/"+ID, help)

	r := g.Group(RoutePrefix)
	r.PUT("", Save)
	r.GET("/:dec", Read)
	r.POST("/:dec", Read)
	r.HEAD("/:dec", Read)
	r.DELETE("/:dec", Delete)
}
