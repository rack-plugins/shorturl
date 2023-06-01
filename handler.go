package shorturl

import (
	"net/http"

	"github.com/fimreal/goutils/ezap"
	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb"
)

func Save(c *gin.Context) {
	var job NewShort
	if err := c.ShouldBind(&job); err != nil {
		ezap.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 记录
	err := DB.Put([]byte(job.Dec), []byte(job.Ori), nil)
	if err != nil {
		ezap.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 读取
	oriUrl, err := DB.Get([]byte(job.Dec), nil)
	if err != nil {
		ezap.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ori": oriUrl, "dec": job.Dec})
}

func Read(c *gin.Context) {
	dec := c.Param("dec")
	oriUrl, err := DB.Get([]byte(dec), nil)
	if err != nil {
		ezap.Error(err.Error())
		if err == leveldb.ErrNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "no record like" + dec})
		} else {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		return
	}
	c.Redirect(http.StatusMovedPermanently, string(oriUrl))
}

func Has(c *gin.Context) {
	dec := c.Param("dec")
	has, err := DB.Has([]byte(dec), nil)
	if err != nil {
		ezap.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"exists": has})
}

func Delete(c *gin.Context) {
	dec := c.Param("dec")
	err := DB.Delete([]byte(dec), nil)
	if err != nil {
		ezap.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "success"})
}
