package main

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
)

func main() {
  fmt.Println("Hello world")

  router := gin.Default()

  router.GET("/hello", func (c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "hello world",})})

  router.Run(":7777")
}
