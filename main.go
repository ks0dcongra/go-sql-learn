package main

import (
  	"database/sql"
  	"fmt"
	"sql_learn/routes"
  _ "github.com/lib/pq"
  	"github.com/gin-gonic/gin"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "postgres"
  dbname   = "test"
)

func main() {
	// connect database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	  "password=%s dbname=%s sslmode=disable",
	  host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	mainServer := gin.Default()
	// 連接Router
	routes.ApiRoutes(mainServer)
	
	// 開啟port
	mainServer.Run(":8090")
}