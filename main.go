package main

import (
    "github.com/gin-gonic/gin"
    "os"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)
var healthy = true

type Customer struct {
    id int 
    email string 
}


func index (c *gin.Context){
    hostname,err := os.Hostname()
    checkErr(err)
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

func fetch (c *gin.Context){
    connStr := os.Getenv("sql_user")+":"+os.Getenv("sql_password")+"@tcp("+os.Getenv("sql_host")+":3306)/"+os.Getenv("sql_db")
    db, err := sql.Open("mysql",connStr)
    checkErr(err)
    defer db.Close()
    cust := new(Customer)
    db.QueryRow("SELECT * FROM customers").Scan(&cust.id,&cust.email)
    checkErr(err)
    c.JSON(200,gin.H{string(cust.id):cust.email})
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func main(){
  app := gin.Default()
  app.GET("/", index)
  app.GET("/healthz", healthz)
  app.GET("/cancer", cancer)
  app.GET("/dbtest",fetch)
  app.Run(":8000")
}

//func homePage(res http.ResponseWriter, req *http.Request) {
	//http.ServeFile(res, req, "index.html")
//}
