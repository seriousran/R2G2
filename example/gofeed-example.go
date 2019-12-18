package main

import (
    "fmt"
    "github.com/mmcdole/gofeed"
)

func main() {
  fp := gofeed.NewParser()
  feed, _ := fp.ParseURL("http://lv99.tistory.com/rss")
  fmt.Println(feed.Title)
}


