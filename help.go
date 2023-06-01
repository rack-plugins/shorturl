package shorturl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func help(ctx *gin.Context) {
	ctx.String(http.StatusOK, `GET/POST/HEAD `+RoutePrefix+`/<dec> => origin url
PUT `+RoutePrefix+` json/xml/form, ori: origin url, dec: short code
DELETE `+RoutePrefix+`/<dec> => delete record
`)
}
