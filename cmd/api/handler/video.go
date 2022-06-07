package handler

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

type req struct {
	Filename string `form:"path"`
}

func Video(c *gin.Context) {
	var r req
	if err := c.ShouldBindQuery(&r); err != nil {
		c.Status(http.StatusBadRequest)
	}

	fileName := os.Getenv("FILEPATH") + r.Filename
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		c.Status(http.StatusNotFound)
	}
	_, err = c.Writer.Write(file)
	if err != nil {
		c.Status(http.StatusNotFound)
	}
	c.Status(http.StatusOK)
}
