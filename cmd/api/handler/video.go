package handler

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type req struct {
	Filename string `form:"path"`
}

func Video(c *gin.Context) {
	var r req
	if err := c.ShouldBindQuery(&r); err != nil {
		c.Status(http.StatusBadRequest)
	}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		c.Status(http.StatusNotFound)
	}
	_, err = c.Writer.Write(file)
	if err != nil {
		c.Status(http.StatusNotFound)
	}
	c.Status(http.StatusOK)
}
