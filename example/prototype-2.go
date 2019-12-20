package main

import (
  "fmt"
  //"log"
  //"encoding/json"
  "github.com/mmcdole/gofeed"
  "github.com/savsgio/atreugo"
)

func main() {
  
  fp := gofeed.NewParser()

  config := &atreugo.Config{
    Addr: "0.0.0.0:8080",
  }
  server := atreugo.New(config)
  
  postRssReq(server, fp)
  
}

type RssRes struct {
  Url   string
  RssBody string
}

func postRssReq(server *atreugo.Atreugo, fp *gofeed.Parser) {
  server.Path("POST", "/rss", func(ctx *atreugo.RequestCtx) error {
    url := string(ctx.FormValue("url"))
    
    feed, _ := fp.ParseURL(url)
    return ctx.JSONResponse(feed)
    
	//rss_res := []RssRes{
	//	{url, url},
	//}
    
    //res, err := json.Marshal(feed)
    //if err != nil {
    //  log.Fatal(err)
    //}
    //return ctx.JSONResponse(res)
  })

  err := server.ListenAndServe()
  if err != nil {
    panic(err)
  }
}