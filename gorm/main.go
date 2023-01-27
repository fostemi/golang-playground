
package main

import (
  // "net/http"
  "github.com/gin-gonic/gin"

  "github.com/fostemi/go-learning/gorm/models"
)

func main() {
  r := gin.Default()

  // r.GET("/", func(c *gin.Context) {
  //   c.JSON(http.StatusOK, gin.H{"data": "hello world"})
  // })

  models.ConnectDatabase()

  r.Run()
}

