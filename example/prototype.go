package main

import (
  "fmt"
  "github.com/mmcdole/gofeed"
  "github.com/gin-gonic/gin"
)

func main() {
  
  fp := gofeed.NewParser()
  
  r := gin.Default()
  
  // example code for GET
  r.GET("/ping", func(c *gin.Context) {
    
    feed, _ := fp.ParseURL("http://lv99.tistory.com/rss")
    fmt.Println(feed.Title)
    
    c.JSON(200, gin.H{
      "message": feed.Title,
    })
  })
  
  // example code for POST
  r.POST("/post", func(c *gin.Context) {
    id := c.Query("id")
    page := c.DefaultQuery("page", "0")
    name := c.PostForm("name")
    message := c.PostForm("message")
    
    fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
  })
  
  // GET query
  // input: url
  // return: parsed rss
  r.GET("/get-rss", func(c *gin.Context) {
    _url := c.Query("url") // shortcut for c.Request.URL.Query().Get("lastname")
    
    fmt.Println(_url)
    
    feed, _ := fp.ParseURL(_url)
    
    fmt.Println(feed)
    
    c.JSON(200, gin.H{
      "message": feed,
    })
    //c.String(http.StatusOK, feed)
  })
  
  r.Run()
}
