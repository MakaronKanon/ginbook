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
  router.GET("/", indexHandler)
  router.GET("/:name", nameHandler)

  router.Run(":7777")
}

func indexHandler(c *gin.Context) {
  c.JSON(200, gin.H{"hej": "yo"})
}

func nameHandler(c *gin.Context) {
  name := c.Param("name")
  name2 := c.Query("name2")

  type msg struct {
    Name string `json:"name"`
    name2 string
  }
  myMsg := msg{Name: name, name2: name2}

  c.JSON(200, myMsg)
}
