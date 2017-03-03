package main

import (
    "github.com/gin-gonic/gin"
    "os"
)
var healthy = true

func index (c *gin.Context){
    hostname,err := os.Hostname()
    if err!=nil {
       c.String(500,"Error")
    }
    c.String(200,hostname)
}

func healthz (c *gin.Context){
    if healthy==true {
     c.String(200,"OK")
    }
}

func cancer (c *gin.Context){
     healthy = false
     c.String(500,"NOT_OK")
}


func main(){
  app := gin.Default()
  app.GET("/", index)
  app.GET("/healthz", healthz)
  app.GET("/cancer", cancer)
  app.Run(":8000")
}
